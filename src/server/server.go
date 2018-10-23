package main

import (
	"github.com/gorilla/mux"
	"gobi/src/controllers"
	"log"
	"net/http"
	"os"
)


func main() {

	// todo:server needs a connection pool ?
	// todo:server add kafka
	// todo:server add influxdb
	// todo:split to smaller services

	router := mux.NewRouter()
	os.Setenv("PORT", ":8000")

	router.HandleFunc("/token", controllers.CreateUserToken).Methods("POST")
	router.HandleFunc("/authenticate", controllers.GetUserToken).Methods("POST")

	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")


	/**
		router.HandleFunc("/organization", controllers.GetOrganizations).Methods("GET")
		router.HandleFunc("/organization/{organization_id}", controllers.GetOrganization).Methods("GET")
		router.HandleFunc("/organization", controllers.CreateOrganization).Methods("POST")
		router.HandleFunc("/organization/{organization_id}", controllers.DeleteOrganization).Methods("DELETE")
		router.HandleFunc("/organization/{organization_id}", controllers.UpdateOrganization).Methods("PUT")

	    router.HandleFunc("/organization/suborganization/{organization_id}", controllers.GetSubOrganizationsByOrganization).Methods("GET")
	    router.HandleFunc("/organization/suborganization/{organization_id}", controllers.AddSubOrganizationsToOrganization).Methods("POST")
	    router.HandleFunc("/organization/suborganization/{organization_id}", controllers.RemoveAllSubOrganizationsFromOrganization).Methods("DELETE")
	    router.HandleFunc("/organization/{organization_id}/suborganization/{sub_organization_id}", controllers.RemoveSubOrganizationFromOrganization).Methods("DELETE")
	    router.HandleFunc("/organization/{organization_id}/suborganization/{sub_organization_id}", controllers.UpdateSubOrganizationForOrganization).Methods("PUT")

	    router.HandleFunc("/organization/building/{id}", controllers.GetBuildingByOrganization).Methods("GET")
	    router.HandleFunc("/organization/building/{organization_id}", controllers.AddBuildingToOrganization).Methods("POST")
	    router.HandleFunc("/organization/building/{organization_id}", controllers.RemoveAllBuildingsFromOrganization).Methods("DELETE")
	    router.HandleFunc("/organization/{organization_id}/building/{building_id}", controllers.RemoveBuildingFromOrganization).Methods("DELETE")
	    router.HandleFunc("/organization/{organization_id}/building/{building_id}", controllers.UpdateBuildingForOrganization).Methods("PUT")

		router.HandleFunc("/suborganizations", controllers.GetSubOrganizations).Methods("GET")
		router.HandleFunc("/suborganization/{sub_organization_id}", controllers.GetSubOrganization).Methods("GET")
		router.HandleFunc("/suborganization", controllers.CreateSubOrganization).Methods("POST")
		router.HandleFunc("/suborganization/{sub_organization_id}", controllers.DeleteSubOrganization).Methods("DELETE")
		router.HandleFunc("/suborganization/{sub_organization_id}", controllers.UpdateSubOrganization).Methods("PUT")

	    router.HandleFunc("/suborganization/building/{sub_organization_id}", controllers.GetBuildingBySubOrganization).Methods("GET")
	    router.HandleFunc("/suborganization/building/{sub_organization_id}", controllers.AddBuildingToSubOrganization).Methods("POST")
	    router.HandleFunc("/suborganization/building/{sub_organization_id}", controllers.RemoveAllBuildingsFromSubOrganization).Methods("DELETE")
        router.HandleFunc("/suborganization/{sub_organization_id}/building/{building_id}", controllers.RemoveBuildingFromSubOrganization).Methods("DELETE")
	    router.HandleFunc("/suborganization/{sub_organization_id}/building/{building_id}", controllers.UpdateBuildingForSubOrganization).Methods("PUT")

	    router.HandleFunc("/buildings", controllers.GetBuildings).Methods("GET")
		router.HandleFunc("/building/{building_id}", controllers.GetBuilding).Methods("GET")
		router.HandleFunc("/building", controllers.CreateBuilding).Methods("POST")
		router.HandleFunc("/building/{building_id}", controllers.DeleteBuilding).Methods("DELETE")
		router.HandleFunc("/building/{building_id}", controllers.UpdateBuilding).Methods("PUT")

	    router.HandleFunc("/buildingsystems", controllers.GetBuildingSystems).Methods("GET")
		router.HandleFunc("/buildingsystem/{building_system_id}", controllers.GetBuildingSystem).Methods("GET")
		router.HandleFunc("/buildingsystem", controllers.CreateBuildingSystem).Methods("POST")
		router.HandleFunc("/buildingsystem/{building_system_id}", controllers.DeleteBuildingSystem).Methods("DELETE")
		router.HandleFunc("/buildingsystem/{building_system_id}", controllers.UpdateBuildingSystem).Methods("PUT")


	    router.HandleFunc("/building/floors/{building_id}", controllers.GetFloorsByBuilding).Methods("GET")
	    router.HandleFunc("/building/floors/{building_id}", controllers.AddFloorsToBuilding).Methods("POST")
	   router.HandleFunc("/building/floors/{building_id}", controllers.RemoveAllFloorsFromBuilding).Methods("DELETE")
	   router.HandleFunc("/building/{building_id}/floors/{floor_id}", controllers.RemoveFloorFromBuilding).Methods("DELETE")
	   router.HandleFunc("/building/{building_id}/floors/{floor_id}", controllers.UpdateFloorForBuilding).Methods("PUT")

	   router.HandleFunc("/building/building_system/{building_id}", controllers.GetBuildingSystemsByBuilding).Methods("GET")
	   router.HandleFunc("/building/building_system/{building_id}", controllers.AddBuildingSystemToBuilding).Methods("POST")
	   router.HandleFunc("/building/building_system/{building_id}", controllers.RemoveAllBuildingSystemsFromBuilding).Methods("DELETE")
	   router.HandleFunc("/building/{building_id}/building_system/{building_system_id}", RemoveBuildingSystemFromBuilding).Methods("DELETE")
	   router.HandleFunc("/building/{building_id}/building_system/{building_system_id}", UpdateBuildingSystemForBuilding).Methods("DELETE")

	    router.HandleFunc("/floors", controllers.GetFloors).Methods("GET")
		router.HandleFunc("/floor/{floor_id}", controllers.Floor).Methods("GET")
		router.HandleFunc("/floor", controllers.CreateFloor).Methods("POST")
		router.HandleFunc("/floor/{floor_id}", controllers.DeleteFloor).Methods("DELETE")
		router.HandleFunc("/floor/{floor_id}", controllers.UpdateFloor).Methods("PUT")

	    router.HandleFunc("/floor/rooms/{floor_id}", controllers.GetRoomsByFloor).Methods("GET")
	    router.HandleFunc("/floor/rooms/{floor_id}", controllers.AddRoomToFloor).Methods("POST")
	    router.HandleFunc("/floor/rooms/{floor_id}", controllers.RemoveRoomsByFloor).Methods("DELETE")
	    router.HandleFunc("/floor/{floor_id}/rooms/{room_id}", controllers.RemoveRoomsForFloor).Methods("DELETE")
	    router.HandleFunc("/floor/{floor_id}/rooms/{floor_id}", controllers.UpdateRoomForFloor).Methods("PUT")

	     router.HandleFunc("/floor/building_system/{floor_id}", controllers.GetBuildingSystemsByFloor).Methods("GET")
	     router.HandleFunc("/floor/building_system/{floor_id}", controllers.AddBuildingSystemsToFloor).Methods("POST")
	     router.HandleFunc("/floor/building_system/{floor_id}", controllers.RemoveAllBuildingSystemsByFloor).Methods("DELETE")
	     router.HandleFunc("/floor/{floor_id}/building_system/{building_system_id}", controllers.RemoveBuildingSystemForFloor).Methods("DELETE")
         router.HandleFunc("/floor/{floor_id}/building_system/{building_system_id}", controllers.UpdateBuildingSystemForFloor).Methods("PUT")

         router.HandleFunc("/floor/gateway/{floor_id}", controllers.GetGateWayByFloor).Methods("GET")
	     router.HandleFunc("/floor/gateway/{floor_id}", controllers.AddGateWayToFloor).Methods("POST")
	     router.HandleFunc("/floor/gateway/{floor_id}", controllers.RemoveAllGateWaysByFloor).Methods("DELETE")
	     router.HandleFunc("/floor/{floor_id}/gateway/{gateway_id}", controllers.RemoveGateWayForFloor).Methods("DELETE")
         router.HandleFunc("/floor/{floor_id}/gateway/{gateway_id}", controllers.UpdateGateWayForFloor).Methods("PUT")


	    router.HandleFunc("/rooms", controllers.GetSubOrganizations).Methods("GET")
		router.HandleFunc("/room/{room_id}", controllers.GetSubOrganization).Methods("GET")
		router.HandleFunc("/room", controllers.CreateSubOrganization).Methods("POST")
		router.HandleFunc("/room/{room_id}", controllers.DeleteSubOrganization).Methods("DELETE")
		router.HandleFunc("/room/{room_id}", controllers.UpdateSubOrganization).Methods("PUT")

         router.HandleFunc("/room/gateway/{floor_id}", controllers.GetGateWayByRoom).Methods("GET")
	     router.HandleFunc("/room/gateway/{floor_id}", controllers.AddGateWayToRoom).Methods("POST")
	     router.HandleFunc("/room/gateway/{floor_id}", controllers.RemoveAllGateWaysByRoom).Methods("DELETE")
	     router.HandleFunc("/room/{room_id}/gateway/{gateway_id}", controllers.RemoveGateWayForRoom).Methods("DELETE")
         router.HandleFunc("/room/{room_id}/gateway/{gateway_id}", controllers.UpdateGateWayForRoom).Methods("PUT")

	    router.HandleFunc("/gateways", controllers.GetGateways).Methods("GET")
		router.HandleFunc("/gateway/{gateway_id}", controllers.GetGetateway).Methods("GET")
		router.HandleFunc("/gateway", controllers.CreateGateways).Methods("POST")
		router.HandleFunc("/gateway/{gateway_id}", controllers.DeleteGateway).Methods("DELETE")
		router.HandleFunc("/gateway/{gateway_id}", controllers.UpdateGateway).Methods("PUT")

	     router.HandleFunc("/gateway/protocol/{gateway_id}", controllers.GetProtocolByGateway).Methods("GET")
	     router.HandleFunc("/gateway/protocol/{gateway_id}", controllers.AddProtocolToGateway).Methods("POST")
	     router.HandleFunc("/gateway/protocol/{gateway_id}", controllers.RemoveAllProtocolsByGateway).Methods("DELETE")
	     router.HandleFunc("/gateway/{gateway_id}/protocol/{protocol_id}", controllers.RemoveProtocolForGateway).Methods("DELETE")
         router.HandleFunc("/gateway/{gateway_id}/protocol/{protocol_id}", controllers.UpdateProtocolForGateway).Methods("PUT")

	     router.HandleFunc("/gateway/device/{gateway_id}", controllers.GetDeviceByGateway).Methods("GET")
	     router.HandleFunc("/gateway/device/{gateway_id}", controllers.AddDeviceToGateway).Methods("POST")
	     router.HandleFunc("/gateway/device/{gateway_id}", controllers.RemoveAllDevicesByGateway).Methods("DELETE")
	     router.HandleFunc("/gateway/{gateway_id}/device/{device_id}", controllers.RemoveDeviceForGateway).Methods("DELETE")
         router.HandleFunc("/gateway/{gateway_id}/device/{device_id}", controllers.UpdateDeviceForGateway).Methods("PUT")

        router.HandleFunc("/devices", controllers.GetDevices).Methods("GET")
		router.HandleFunc("/device/{device_id}", controllers.GetDevice).Methods("GET")
		router.HandleFunc("/device", controllers.CreateDevice).Methods("POST")
		router.HandleFunc("/device/{device_id}", controllers.DeleteDevice).Methods("DELETE")
		router.HandleFunc("/device/{device_id}", controllers.UpdateDevice).Methods("PUT")

	     router.HandleFunc("/device/protocol/{device_id}", controllers.GetProtocolByDevice).Methods("GET")
	     router.HandleFunc("/device/protocol/{device_id}", controllers.AddProtocolToDevice).Methods("POST")
	     router.HandleFunc("/device/protocol/{device_id}", controllers.RemoveAllProtocolsByDevice).Methods("DELETE")
	     router.HandleFunc("/device/{device_id}/protocol/{protocol_id}", controllers.RemoveProtocolForDevice).Methods("DELETE")
         router.HandleFunc("/device/{device_id}/protocol/{protocol_id}", controllers.UpdateProtocolForDevice).Methods("PUT")

	     router.HandleFunc("/device/building_systems/{device_id}", controllers.GetBuildingSystemsByDevice).Methods("GET")
	     router.HandleFunc("/device/building_systems/{device_id}", controllers.AddBuildingSystemToDevice).Methods("POST")
	     router.HandleFunc("/device/building_systems/{device_id}", controllers.RemoveAllBuildingSystemByDevice).Methods("DELETE")
	     router.HandleFunc("/device/{device_id}/device/{building_system_id}", controllers.RemoveBuildingSystemForDevice).Methods("DELETE")
         router.HandleFunc("/device/{device_id}/device/{building_system_id}", controllers.UpdateBuildingSystemForDevice).Methods("PUT")

         router.HandleFunc("/device/sensor/{device_id}", controllers.GetSensorsByDevice).Methods("GET")
	     router.HandleFunc("/device/sensor/{device_id}", controllers.AddSensorsToDevice).Methods("POST")
	     router.HandleFunc("/device/sensor/{device_id}", controllers.RemoveAllSensorsByDevice).Methods("DELETE")
	     router.HandleFunc("/device/{device_id}/sensor/{sensor_id}", controllers.RemoveSensorForDevice).Methods("DELETE")
         router.HandleFunc("/device/{device_id}/sensor/{sensor_id}", controllers.UpdateSensorForDevice).Methods("PUT")


        router.HandleFunc("/sensors", controllers.GetSensors).Methods("GET")
		router.HandleFunc("/sensor/{sensor_id}", controllers.GetSensor).Methods("GET")
		router.HandleFunc("/sensor", controllers.CreateSensor).Methods("POST")
		router.HandleFunc("/sensor/{sensor_id}", controllers.DeleteSensor).Methods("DELETE")
		router.HandleFunc("/sensor/{sensor_id}", controllers.UpdateSensor).Methods("PUT")

	     router.HandleFunc("/sensor/protocol/{sensor_id}", controllers.GetProtocolBySensor).Methods("GET")
	     router.HandleFunc("/sensor/protocol/{sensor_id}", controllers.AddProtocolToSensor).Methods("POST")
	     router.HandleFunc("/sensor/protocol/{sensor_id}", controllers.RemoveAllProtocolsBySensor).Methods("DELETE")
	     router.HandleFunc("/sensor/{sensor_id}/protocol/{protocol_id}", controllers.RemoveProtocolForSensor).Methods("DELETE")
         router.HandleFunc("/sensor/{sensor_id}/protocol/{protocol_id}", controllers.UpdateProtocolForSensor).Methods("PUT")

	     router.HandleFunc("/sensor/building_systems/{sensor_id}", controllers.GetBuildingSystemsBySensor).Methods("GET")
	     router.HandleFunc("/sensor/building_systems/{sensor_id}", controllers.AddBuildingSystemToSensor).Methods("POST")
	     router.HandleFunc("/sensor/building_systems/{sensor_id}", controllers.RemoveAllBuildingSystemBySensor).Methods("DELETE")
	     router.HandleFunc("/sensor/{sensor_id}/device/{building_system_id}", controllers.BuildingSystemForSensor).Methods("DELETE")
         router.HandleFunc("/sensor/{sensor_id}/device/{building_system_id}", controllers.UpdateBuildingSystemForSensor).Methods("PUT")

	    router.HandleFunc("/protocol", controllers.GetProtocols).Methods("GET")
		router.HandleFunc("/protocol/{protocol_id}", controllers.GetProtocol).Methods("GET")
		router.HandleFunc("/protocol", controllers.CreateProtocol).Methods("POST")
		router.HandleFunc("/protocol/{protocol_id}", controllers.DeleteProtocol).Methods("DELETE")
		router.HandleFunc("/protocol/{protocol_id}", controllers.UpdateProtocol).Methods("PUT")

	    router.HandleFunc("/observable", controllers.GetObservables).Methods("GET")
		router.HandleFunc("/observable/{observable_id}", controllers.GetObservable).Methods("GET")
		router.HandleFunc("/observable", controllers.CreateObservable).Methods("POST")
		router.HandleFunc("/observable/{observable_id}", controllers.DeleteObservable).Methods("DELETE")
		router.HandleFunc("/observable/{observable_id}", controllers.UpdateObservable).Methods("PUT")


	 */
	log.Fatal(http.ListenAndServe( ":8000", router))
}
