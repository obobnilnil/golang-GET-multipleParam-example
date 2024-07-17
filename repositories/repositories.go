package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"multipleParam_git/models"

	"github.com/lib/pq"
)

type RepositoryPort interface {
	GetUniversalInfoRepositories2(catalog []string) ([]models.UniversalInfo, error)
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) GetUniversalInfoRepositories2(catalogs []string) ([]models.UniversalInfo, error) {
	var query string
	var rows *sql.Rows
	var err error

	// Check if 'All' is in the catalogs slice
	if len(catalogs) == 1 && catalogs[0] == "All" {
		// If 'All' is the only item, select all records
		query = `
            SELECT uin_pk, uin_catalog_lc, uin_abbreviation_lc, uin_desc_lc, uin_catalog_en, uin_abbreviation_en, uin_desc_en, uin_info, uin_ref
            FROM universal_info`
		rows, err = r.db.Query(query)
	} else {
		// Otherwise, use the catalogs provided to filter results
		query = `
            SELECT uin_pk, uin_catalog_lc, uin_abbreviation_lc, uin_desc_lc, uin_catalog_en, uin_abbreviation_en, uin_desc_en, uin_info, uin_ref
            FROM universal_info
            WHERE uin_catalog_en = ANY($1)`
		rows, err = r.db.Query(query, pq.Array(catalogs))
	}

	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	var data []models.UniversalInfo
	for rows.Next() {
		var info models.UniversalInfo
		err := rows.Scan(&info.UinPK, &info.UinCatalogLC, &info.UinAbbreviationLC, &info.UinDescLC, &info.UinCatalogEN, &info.UinAbbreviationEN, &info.UinDescEN, &info.UinInfo, &info.UinRef)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		data = append(data, info)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no records found for catalogs: %v", catalogs)
	}

	return data, nil
}
