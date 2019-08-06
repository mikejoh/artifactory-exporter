package collector

type BasicCredentials struct {
	Username string
	Password string
}

type storageInfoResponse struct {
	StorageSummary struct {
		BinariesSummary struct {
			BinariesCount  float64 `json:"binariesCount,string"`
			BinariesSize   string `json:"binariesSize"`
			ArtifactsSize  string `json:"artifactsSize"`
			Optimization   string `json:"optimization"`
			ItemsCount     float64 `json:"itemsCount,string"`
			ArtifactsCount float64 `json:"artifactsCount,string"`
		} `json:"binariesSummary"`
		FileStoreSummary struct {
			StorageType      string `json:"storageType"`
			StorageDirectory string `json:"storageDirectory"`
			TotalSpace       string `json:"totalSpace"`
			UsedSpace        string `json:"usedSpace"`
			FreeSpace        string `json:"freeSpace"`
		} `json:"fileStoreSummary"`
		RepositoriesSummaryList []struct {
			RepoKey      string `json:"repoKey"`
			RepoType     string `json:"repoType"`
			FoldersCount int    `json:"foldersCount"`
			FilesCount   int    `json:"filesCount"`
			UsedSpace    string `json:"usedSpace"`
			ItemsCount   int    `json:"itemsCount"`
			PackageType  string `json:"packageType,omitempty"`
			Percentage   string `json:"percentage,omitempty"`
		} `json:"repositoriesSummaryList"`
	} `json:"storageSummary"`
	FileStoreSummary struct {
		StorageType      string `json:"storageType"`
		StorageDirectory string `json:"storageDirectory"`
		TotalSpace       string `json:"totalSpace"`
		UsedSpace        string `json:"usedSpace"`
		FreeSpace        string `json:"freeSpace"`
	} `json:"fileStoreSummary"`
	RepositoriesSummaryList []struct {
		RepoKey      string `json:"repoKey"`
		RepoType     string `json:"repoType"`
		FoldersCount int    `json:"foldersCount"`
		FilesCount   int    `json:"filesCount"`
		UsedSpace    string `json:"usedSpace"`
		ItemsCount   int    `json:"itemsCount"`
		PackageType  string `json:"packageType,omitempty"`
		Percentage   string `json:"percentage,omitempty"`
	} `json:"repositoriesSummaryList"`
	BinariesSummary struct {
		BinariesCount  float64 `json:"binariesCount,string"`
		BinariesSize   string `json:"binariesSize"`
		ArtifactsSize  string `json:"artifactsSize"`
		Optimization   string `json:"optimization"`
		ItemsCount     float64 `json:",string"`
		ArtifactsCount float64 `json:"artifactsCount,string"`
	} `json:"binariesSummary"`
}