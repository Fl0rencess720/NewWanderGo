package Config

import (
	"github.com/go-redis/redis"
)

var buildingGeoLoc = []*redis.GeoLocation{
	{Name: "Library", Longitude: 115.799759, Latitude: 28.656642},
	{Name: "ShuRenSqure", Longitude: 115.803208, Latitude: 28.65666},
	{Name: "JiChuShiYanBuilding", Longitude: 115.797922, Latitude: 28.657403},
	{Name: "LiShengBuilding", Longitude: 115.798373, Latitude: 28.658961},
	{Name: "XianSuGarden", Longitude: 115.799698, Latitude: 28.65874},
	{Name: "ZhiHuaKeJiBuilding", Longitude: 115.800395, Latitude: 28.659441},
	{Name: "ZhiHuaJingGuanBuilding", Longitude: 115.801125, Latitude: 28.659587},
	{Name: "XinGongBuilding", Longitude: 115.799666, Latitude: 28.661154},
	{Name: "JiDianBuilding", Longitude: 115.800352, Latitude: 28.661945},
	{Name: "QiCheDianZiBuilding", Longitude: 115.800604, Latitude: 28.662524},
	{Name: "JianGongBuilding", Longitude: 115.801436, Latitude: 28.662868},
	{Name: "HuiYuanBuilding", Longitude: 115.804284, Latitude: 28.663725},
	{Name: "DiningHallOne", Longitude: 115.8043, Latitude: 28.664388},
	{Name: "CommercialStreet", Longitude: 115.806307, Latitude: 28.665066},
	{Name: "XiuXianSqure", Longitude: 115.807868, Latitude: 28.665819},
	{Name: "XiuXianTrack", Longitude: 115.810137, Latitude: 28.665433},
	{Name: "Natatorium", Longitude: 115.811124, Latitude: 28.664209},
	{Name: "BaiFanField", Longitude: 115.811805, Latitude: 28.661752},
	{Name: "Gymnasium", Longitude: 115.811607, Latitude: 28.659257},
	{Name: "ZhengQiSqure", Longitude: 115.806291, Latitude: 28.659173},
	{Name: "WenFaBuilding", Longitude: 115.804467, Latitude: 28.66018},
	{Name: "YiShuBuilding", Longitude: 115.80805, Latitude: 28.660886},
	{Name: "Hospital", Longitude: 115.809359, Latitude: 28.663993},
	{Name: "TianJianTrack", Longitude: 115.79612, Latitude: 28.653882},
	{Name: "ChangHaiBuilding", Longitude: 115.797252, Latitude: 28.651514},
}

func AddGeoInfo() {
	for i := range buildingGeoLoc {
		GLOBAL_RDB.GeoAdd("NCU:Buildings", buildingGeoLoc[i])
	}
}
