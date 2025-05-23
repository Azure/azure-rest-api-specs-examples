package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/OperationList.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationListResult = armpostgresqlflexibleservers.OperationListResult{
		// 	Value: []*armpostgresqlflexibleservers.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/queryTexts/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/queryTexts/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/recommendations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of recommended indexes for the Azure Database for PostgreSQL Flexible Server"),
		// 				Operation: to.Ptr("List Azure Database for PostgreSQL Flexible Server recommended indexes for tuning."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Recommended Indexes"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Tuning Options available for the Azure Database for PostgreSQL Flexible Server"),
		// 				Operation: to.Ptr("List Azure Database for PostgreSQL Flexible Server Tuning Options."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Tuning Option"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns a tuning option"),
		// 				Operation: to.Ptr("Get a single tuning option"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Tuning Option"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a migration with the specified parameters."),
		// 				Operation: to.Ptr("Create a database migration workflow."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Migration Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the properties for the specified migration workflow."),
		// 				Operation: to.Ptr("Get Migration Workflow details."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Migration Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("List of migration workflows for the specified database server."),
		// 				Operation: to.Ptr("List Migration Workflow details."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Migration Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Update the properties for the specified migration."),
		// 				Operation: to.Ptr("Update a database migration workflow."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Migration Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/migrations/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing migration workflow."),
		// 				Operation: to.Ptr("Delete Migration Workflow."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Migration Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/getAutoMigrationFreeSlots/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of free / available slots for auto migration of PostgreSQL server"),
		// 				Operation: to.Ptr("Get auto migration free / available slots for PostgreSQL"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Auto migration workflow to get free / available slots for auto migration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/getLatestAutoMigrationSchedule/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the instance of the latest auto migration schedule for PostgreSQL server"),
		// 				Operation: to.Ptr("Get latest auto migration schedule for PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Auto migration workflow to get latest auto migration schedule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/updateAutoMigrationSchedule/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Update auto migration schedule for the PostgreSQL server"),
		// 				Operation: to.Ptr("Update auto migration schedule for the PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Auto migration workflow to update auto migration schedule"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/startSession/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Starts a server configuration tuning session on a server"),
		// 				Operation: to.Ptr("Start tuning configuration session"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/stopSession/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Stops the server configuration tuning session on a server"),
		// 				Operation: to.Ptr("Stop tuning configuration session"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/sessions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the list of server configuration tuning sessions on a server"),
		// 				Operation: to.Ptr("Server configuration tuning sessions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/enable/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Enables server configuration tuning feature on the server"),
		// 				Operation: to.Ptr("Enable server configuration tuning feature"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/startSession/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Starts a server configuration tuning session on a server"),
		// 				Operation: to.Ptr("Start tuning configuration session"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/stopSession/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Stops the server configuration tuning session on a server"),
		// 				Operation: to.Ptr("Stop tuning configuration session"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/sessions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the list of server configuration tuning sessions on a server"),
		// 				Operation: to.Ptr("Server configuration tuning sessions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/sessionDetails/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the list of iterations for a specified server configuration tuning session on a server"),
		// 				Operation: to.Ptr("Server configuration session details"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/disable/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Disables server configuration tuning feature on the server"),
		// 				Operation: to.Ptr("Disable server configuration tuning feature"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Configuration Tuning"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/restart/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Restarts an existing server"),
		// 				Operation: to.Ptr("Restart PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/start/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Starts an existing server"),
		// 				Operation: to.Ptr("Start PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/stop/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Stops an existing server"),
		// 				Operation: to.Ptr("Stops an existing server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/replicas/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of servers or gets the properties for the specified server."),
		// 				Operation: to.Ptr("List/Get PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server with the specified parameters or update the properties or tags for the specified server."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server with the specified parameters or update the properties or tags for the specified server."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing server."),
		// 				Operation: to.Ptr("Delete PostgreSQL server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/capabilities/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the capabilities for this subscription in a given location"),
		// 				Operation: to.Ptr("Gets the capabilities for this subscription"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Location Capability"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/capabilities/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the capabilities for this subscription in a given location"),
		// 				Operation: to.Ptr("Gets the capabilities for this subscription"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Location Capability"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/capabilities/{serverName}/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the capabilities for this subscription in a given location"),
		// 				Operation: to.Ptr("Gets the capabilities for this subscription"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Location Capability"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Flexible server database backups"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Flexible server database backups"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/backups/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Flexible server database backups"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advisors/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/waitStatistics/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/queryStatistics/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/topQueryStatistics/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/resetQueryPerformanceInsightData/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advisors/recommendedActions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/checkMigrationNameAvailability/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Checks the availability of the given migration name."),
		// 				Operation: to.Ptr("Check the availability of the given migration name."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Migration Name Availability Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/assessForMigration/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Performs a migration assessment with the specified parameters"),
		// 				Operation: to.Ptr("Get Migration Assessment"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Migration Assessments"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of server administrators or gets the properties for the specified server administrator."),
		// 				Operation: to.Ptr("List/Get PostgreSQL server administrator"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Administrator"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of server administrators or gets the properties for the specified server administrator."),
		// 				Operation: to.Ptr("List/Get PostgreSQL server administrator"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Administrator"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing PostgreSQL server administrator."),
		// 				Operation: to.Ptr("Delete PostgreSQL server administrator"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Administrator"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server administrator with the specified parameters or update the properties or tags for the specified server administrator."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL server administrator"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Administrator"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server administrator with the specified parameters or update the properties or tags for the specified server administrator."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL server administrator"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Administrator"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a firewall rule with the specified parameters or update an existing rule."),
		// 				Operation: to.Ptr("Create/Update Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of firewall rules for a server or gets the properties for the specified firewall rule."),
		// 				Operation: to.Ptr("List/Get Firewall Rules"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/firewallRules/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing firewall rule."),
		// 				Operation: to.Ptr("Delete Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/databases/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of  PostgreSQL server databases or gets the database for the specified server."),
		// 				Operation: to.Ptr("List/Get PostgreSQL server configuration"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/databases/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or Updates the database of a PostgreSQL server."),
		// 				Operation: to.Ptr("Creates/Updates PostgreSQL server database"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/databases/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Delete the database of a PostgreSQL server"),
		// 				Operation: to.Ptr("Delete PostgreSQL server database"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Database"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/configurations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of  PostgreSQL server configurations or gets the configurations for the specified server."),
		// 				Operation: to.Ptr("List/Get PostgreSQL server configuration"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/configurations/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Updates the configuration of a PostgreSQL server."),
		// 				Operation: to.Ptr("Update PostgreSQL server configuration"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/getSourceDatabaseList/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Fetch Databases Pre Migration Workflow"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualendpoints/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or Updates VirtualEndpoint"),
		// 				Operation: to.Ptr("VirtualEndpoint Create or Update operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Virtual Endpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualendpoints/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Patches the VirtualEndpoint. Currently patch does a full replace"),
		// 				Operation: to.Ptr("VirtualEndpoint Patch operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Virtual Endpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualendpoints/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes the VirtualEndpoint"),
		// 				Operation: to.Ptr("VirtualEndpoint Delete operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Virtual Endpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualendpoints/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the VirtualEndpoint details"),
		// 				Operation: to.Ptr("VirtualEndpoint Get operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Virtual Endpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/virtualendpoints/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Lists the VirtualEndpoints"),
		// 				Operation: to.Ptr("VirtualEndpoint List operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Virtual Endpoint"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/testConnectivity/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Pre-migration Workflow for checking the server connectivity"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/logFiles/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return a list of server log Files for a PostgreSQL Flexible server with File download links"),
		// 				Operation: to.Ptr("List/Get PostgreSQL Flexible Server Log Files"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Flexible Server Log Files"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/startLtrBackup/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Start long term backup for a server"),
		// 				Operation: to.Ptr("Start long term backup for a server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Long term backup operation."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/ltrPreBackup/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Checks if a server is ready for a long term backup"),
		// 				Operation: to.Ptr("Checks if a server is ready for a long term backup"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Long term backup operation."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/ltrBackupOperations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the PostgreSQL server long term backup operation tracking by backup name."),
		// 				Operation: to.Ptr("Get PostgreSQL server long term backup operation by backup name"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Long term backup operation."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/ltrBackupOperations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of  PostgreSQL server long term backup operation tracking."),
		// 				Operation: to.Ptr("List/Get PostgreSQL server long term backup operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Long term backup operation."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Advanced Threat Protection or gets the properties for the specified Advanced Threat Protection."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL Flexible Server Advanced Threat Protection states."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Server Advanced Threat Protection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Advanced Threat Protection or gets the properties for the specified Advanced Threat Protection."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL Flexible Server Advanced Threat Protection states."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Server Advanced Threat Protection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/advancedThreatProtectionSettings/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Enables/Disables Azure Database for PostgreSQL Flexible Server Advanced Threat Protection"),
		// 				Operation: to.Ptr("Update Azure Database for PostgreSQL Server Advanced Threat Protection state"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Server Advanced Threat Protection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/resourceType/usages/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the quota usages of a subscription"),
		// 				Operation: to.Ptr("Gets quota usage"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnectionProxies/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connection proxies or gets the properties for the specified private endpoint connection proxy."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL Flexible Server Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnectionProxies/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection proxy resource."),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnectionProxies/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a private endpoint connection proxy with the specified parameters or updates the properties or tags for the specified private endpoint connection proxy"),
		// 				Operation: to.Ptr("Create/Update Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Validates a private endpoint connection create call from NRP side"),
		// 				Operation: to.Ptr("Validate Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Creation by NRP"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connections or gets the properties for the specified private endpoint connection."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection"),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnections/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Approves or rejects an existing private endpoint connection"),
		// 				Operation: to.Ptr("Approve or Reject Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL - Flexible Server Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateLinkResources/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return a list containing private link resource or gets the specified private link resource."),
		// 				Operation: to.Ptr("List/Get PostgreSQL private link resource"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private link resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnectionProxies/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connections or gets the properties for the specified private endpoint connection via proxy"),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL SGv2 Private Endpoint Connection Via Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnectionProxies/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a private endpoint connection with the specified parameters or updates the properties or tags for the specified private endpoint connection via proxy"),
		// 				Operation: to.Ptr("Create/Update Azure Database for PostgreSQL SGv2 Private Endpoint Connection Via Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnectionProxies/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection via proxy"),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL SGv2 Private Endpoint Connection Via Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Validates a private endpoint connection creation by NRP"),
		// 				Operation: to.Ptr("Validate Azure Database for PostgreSQL SGv2 Private Endpoint Connection Creation by NRP"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connections or gets the properties for the specified private endpoint connection"),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL SGv2 Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Approves or rejects an existing private endpoint connection"),
		// 				Operation: to.Ptr("Approve or Reject Azure Database for PostgreSQL SGv2 Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection"),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL SGv2 Private Endpoint Connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateLinkResources/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Get the private link resources for the corresponding PostgreSQL SGv2"),
		// 				Operation: to.Ptr("Get the private link resources for the corresponding PostgreSQL SGv2"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Link Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnectionsApproval/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Determines if user is allowed to approve a private endpoint connection for PostgreSQL SGv2"),
		// 				Operation: to.Ptr("Private Endpoint Connections Approval for PostgreSQL SGv2"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL SGv2 Private Endpoint Connections Approval Resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/securityAlertPolicies/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Retrieve details of the server threat detection policy configured on a given server"),
		// 				Operation: to.Ptr("Get server threat detection policy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Threat Detection Policy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/securityAlertPolicies/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Change the server threat detection policy for a given server"),
		// 				Operation: to.Ptr("Update server threat detection policy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server Threat Detection Policy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/securityAlertPoliciesOperationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of Server threat detection operation result."),
		// 				Operation: to.Ptr("List/Get Server threat detection operation result."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server threat detection operation result"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/keys/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of server keys or gets the properties for the specified server key."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL Server Key(s)"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Server Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/keys/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a key with the specified parameters or update the properties or tags for the specified server key."),
		// 				Operation: to.Ptr("Create/Update Azure Database for PostgreSQL Server Keys"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Server Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/keys/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing server key."),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL Server Key"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Server Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/serverKeyAzureAsyncOperation/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets in-progress operations on data encryption server keys"),
		// 				Operation: to.Ptr("Data Encryption server keys operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Data Encryption server key operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/serverKeyOperationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets in-progress operations on data encryption server keys"),
		// 				Operation: to.Ptr("Data Encryption server keys operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Data Encryption server key operation"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/advisors/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of advisors"),
		// 				Operation: to.Ptr("Return the list of advisors"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Advisors"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/advisors/recommendedActions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of recommended actions"),
		// 				Operation: to.Ptr("Return the list of recommended actions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Recommended Actions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/advisors/recommendedActionSessions/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Make recommendations"),
		// 				Operation: to.Ptr("Make recommendations"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Recommended Actions"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/queryTexts/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the text of a query"),
		// 				Operation: to.Ptr("List query text of a query"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Query Texts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/queryTexts/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the texts for a list of queries"),
		// 				Operation: to.Ptr("List query texts for a list of queries"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Query Texts"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/resetQueryPerformanceInsightData/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Reset Query Performance Insight data"),
		// 				Operation: to.Ptr("Reset Query Performance Insight data"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Reset Query Performance Insight data"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/topQueryStatistics/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of Query Statistics for the top queries."),
		// 				Operation: to.Ptr("List/Get Query Statistic(s) for top queries"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Top Query Statistics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/waitStatistics/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return wait statistics for an instance"),
		// 				Operation: to.Ptr("List Wait Statistics for an instance"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Wait Statistics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/privateEndpointConnectionsApproval/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Determines if user is allowed to approve a private endpoint connection"),
		// 				Operation: to.Ptr("Private Endpoint Connections Approval"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Private Endpoint Connections Approval"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnectionsApproval/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Determines if user is allowed to approve a private endpoint connection"),
		// 				Operation: to.Ptr("Private Endpoint Connections Approval"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Private Endpoint Connections Approval"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateLinkResources/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Get the private link resources for the corresponding PostgreSQL Server"),
		// 				Operation: to.Ptr("Get the private link resources for the corresponding PostgreSQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private link resource"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnections/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connections or gets the properties for the specified private endpoint connection."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private endpoint connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnections/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection"),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private endpoint connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnections/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Approves or rejects an existing private endpoint connection"),
		// 				Operation: to.Ptr("Approve or Reject Azure Database for PostgreSQL private endpoint connection"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private endpoint connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/privateEndpointConnectionOperationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the result for a private endpoint connection operation"),
		// 				Operation: to.Ptr("Get private endpoint connection operation status"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private endpoint connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/privateEndpointConnectionAzureAsyncOperation/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the result for a private endpoint connection operation"),
		// 				Operation: to.Ptr("Get private endpoint connection operation status"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL private endpoint connection"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/privateEndpointConnectionProxyOperationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the result for a private endpoint connection proxy operation"),
		// 				Operation: to.Ptr("Get private endpoint connection proxy operation status"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnectionProxies/validate/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Validates a private endpoint connection create call from NRP side"),
		// 				Operation: to.Ptr("Validate Azure Database for PostgreSQL Private Endpoint Connection Creation by NRP"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnectionProxies/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of private endpoint connection proxies or gets the properties for the specified private endpoint connection proxy."),
		// 				Operation: to.Ptr("List/Get Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnectionProxies/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a private endpoint connection proxy with the specified parameters or updates the properties or tags for the specified private endpoint connection proxy."),
		// 				Operation: to.Ptr("Create/Update Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateEndpointConnectionProxies/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing private endpoint connection proxy"),
		// 				Operation: to.Ptr("Delete Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/privateEndpointConnectionProxyAzureAsyncOperation/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the result for a private endpoint connection proxy operation"),
		// 				Operation: to.Ptr("Get private endpoint connection proxy operation status"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Azure Database for PostgreSQL Private Endpoint Connection Proxy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/performanceTiers/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Performance Tiers available."),
		// 				Operation: to.Ptr("List Performance Tiers"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Performance Tiers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/operationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return ResourceGroup based PostgreSQL Server Operation Results"),
		// 				Operation: to.Ptr("Get PostgreSQL ResourceGroup based Server Operation Results "),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQ ResourceGroup based Server Operation Results"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/recoverableServers/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the recoverable PostgreSQL Server info"),
		// 				Operation: to.Ptr("Get Recoverable PostgreSQL Server info"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Recoverable PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/replicas/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Get read replicas of a PostgreSQL server"),
		// 				Operation: to.Ptr("Get PostgreSQL read replicas"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of servers or gets the properties for the specified server."),
		// 				Operation: to.Ptr("List/Get PostgreSQL Servers"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server with the specified parameters or update the properties or tags for the specified server."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing server."),
		// 				Operation: to.Ptr("Delete PostgreSQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/performanceTiers/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Performance Tiers available."),
		// 				Operation: to.Ptr("List Performance Tiers"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Performance Tiers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/operationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return PostgreSQL Server Operation Results"),
		// 				Operation: to.Ptr("Get PostgreSQL Server Operation Results"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Operation Results"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/restart/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Restarts a specific server."),
		// 				Operation: to.Ptr("Restart PostgreSQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/securityAlertPoliciesAzureAsyncOperation/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of Server threat detection operation result."),
		// 				Operation: to.Ptr("List/Get Server threat detection operation result."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Server threat detection operation result"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/administrators/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets a list of PostgreSQL server administrators."),
		// 				Operation: to.Ptr("Get Administrators of PostgreSQL server."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Administrator of PostgreSQL server."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/administrators/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates PostgreSQL server administrator with the specified parameters."),
		// 				Operation: to.Ptr("Create/Update Administrator of PostgreSQL server."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Administrator of PostgreSQL server."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/administrators/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing administrator of PostgreSQL server."),
		// 				Operation: to.Ptr("Delete Administrator of PostgreSQL server."),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Administrator of PostgreSQL server."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/administratorAzureAsyncOperation/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets in-progress operations on PostgreSQL server administrators"),
		// 				Operation: to.Ptr("PostgreSQL server administrator operation"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Administrator opertiaons of PostgreSQL server."),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/administratorOperationResults/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return PostgreSQL Server administrator operation results"),
		// 				Operation: to.Ptr("Get PostgreSQL server operation results"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL server operation results"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/register/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Register PostgreSQL Resource Provider"),
		// 				Operation: to.Ptr("Register PostgreSQL Resource Provider"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Microsoft Database For PostgreSQL Resource Provider"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/updateConfigurations/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Update configurations for the specified server"),
		// 				Operation: to.Ptr("Batch Update Server Configurations"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/checkNameAvailability/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Verify whether given server name is available for provisioning worldwide for a given subscription."),
		// 				Operation: to.Ptr("Check Server Name Availability"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/configurations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of configurations for a server or gets the properties for the specified configuration."),
		// 				Operation: to.Ptr("List/Get Configurations"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/configurations/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Update the value for the specified configuration"),
		// 				Operation: to.Ptr("Update Configuration"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/virtualNetworkRules/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of virtual network rules or gets the properties for the specified virtual network rule."),
		// 				Operation: to.Ptr("List/Get Virtual Network Rule(s)"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Virtual Network Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/virtualNetworkRules/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a virtual network rule with the specified parameters or update the properties or tags for the specified virtual network rule."),
		// 				Operation: to.Ptr("Create/Update Virtual Network Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Virtual Network Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/virtualNetworkRules/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing Virtual Network Rule"),
		// 				Operation: to.Ptr("Delete Virtual Network Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Virtual Network Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/locations/azureAsyncOperation/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return PostgreSQL Server Operation Results"),
		// 				Operation: to.Ptr("Get PostgreSQL Server Operation Results"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server Operation Results"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of PostgreSQL Databases or gets the properties for the specified Database."),
		// 				Operation: to.Ptr("List/Get PostgreSQL Database"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Databases"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a PostgreSQL Database with the specified parameters or update the properties for the specified Database."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL Database"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Databases"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/databases/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing PostgreSQL Database."),
		// 				Operation: to.Ptr("Delete PostgreSQL Database"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Databases"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/firewallRules/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of firewall rules for a server or gets the properties for the specified firewall rule."),
		// 				Operation: to.Ptr("List/Get Firewall Rules"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/firewallRules/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a firewall rule with the specified parameters or update an existing rule."),
		// 				Operation: to.Ptr("Create/Update Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/firewallRules/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing firewall rule."),
		// 				Operation: to.Ptr("Delete Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/logFiles/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of PostgreSQL LogFiles."),
		// 				Operation: to.Ptr("List/Get PostgreSQL LogFiles"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL LogFiles"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/performanceTiers/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Returns the list of Performance Tiers available."),
		// 				Operation: to.Ptr("List Performance Tiers"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Performance Tiers"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/operations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of PostgreSQL Operations."),
		// 				Operation: to.Ptr("List/Get PostgreSQL Operations"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Operations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return types of metrics that are available for databases"),
		// 				Operation: to.Ptr("Get database metric definitions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 			Properties: map[string]any{
		// 				"serviceSpecification": map[string]any{
		// 					"metricSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "cpu_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "CPU percent",
		// 							"displayName": "CPU percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "memory_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Memory percent",
		// 							"displayName": "Memory percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "io_consumption_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "IO percent",
		// 							"displayName": "IO percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage percent",
		// 							"displayName": "Storage percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_used",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage used",
		// 							"displayName": "Storage used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_limit",
		// 							"aggregationType": "Maximum",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage limit",
		// 							"displayName": "Storage limit",
		// 							"supportedAggregationTypes":[]any{
		// 								"Maximum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "serverlog_storage_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Server Log storage percent",
		// 							"displayName": "Server Log storage percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "serverlog_storage_usage",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Server Log storage used",
		// 							"displayName": "Server Log storage used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "serverlog_storage_limit",
		// 							"aggregationType": "Maximum",
		// 							"category": "Saturation",
		// 							"displayDescription": "Server Log storage limit",
		// 							"displayName": "Server Log storage limit",
		// 							"supportedAggregationTypes":[]any{
		// 								"Maximum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "active_connections",
		// 							"aggregationType": "Average",
		// 							"category": "Traffic",
		// 							"displayDescription": "Active Connections",
		// 							"displayName": "Active Connections",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "connections_failed",
		// 							"aggregationType": "Total",
		// 							"category": "Errors",
		// 							"displayDescription": "Failed Connections",
		// 							"displayName": "Failed Connections",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "backup_storage_used",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Backup Storage Used",
		// 							"displayName": "Backup Storage Used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"supportedTimeGrainTypes":[]any{
		// 								"PT15M",
		// 								"PT30M",
		// 								"PT1H",
		// 								"PT6H",
		// 								"PT12H",
		// 								"P1D",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "network_bytes_egress",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Network Out across active connections",
		// 							"displayName": "Network Out",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "network_bytes_ingress",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Network In across active connections",
		// 							"displayName": "Network In",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "pg_replica_log_delay_in_seconds",
		// 							"aggregationType": "Maximum",
		// 							"category": "Latency",
		// 							"displayDescription": "Replica lag in seconds",
		// 							"displayName": "Replica Lag",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Seconds",
		// 						},
		// 						map[string]any{
		// 							"name": "pg_replica_log_delay_in_bytes",
		// 							"aggregationType": "Maximum",
		// 							"category": "Latency",
		// 							"displayDescription": "Lag in bytes of the most lagging replica",
		// 							"displayName": "Max Lag Across Replicas",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for PostgreSQL servers"),
		// 				Operation: to.Ptr("Read server log definitions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("The log definition of servers"),
		// 			},
		// 			Properties: map[string]any{
		// 				"serviceSpecification": map[string]any{
		// 					"logSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "PostgreSQLLogs",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "PostgreSQL Server Logs",
		// 						},
		// 						map[string]any{
		// 							"name": "QueryStoreRuntimeStatistics",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "PostgreSQL Query Store Runtime Statistics",
		// 						},
		// 						map[string]any{
		// 							"name": "QueryStoreWaitStatistics",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "PostgreSQL Query Store Wait Statistics",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the disagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/servers/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return types of metrics that are available for databases"),
		// 				Operation: to.Ptr("Get database metric definitions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 			Properties: map[string]any{
		// 				"serviceSpecification": map[string]any{
		// 					"metricSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "cpu_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "CPU percent",
		// 							"displayName": "CPU percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "memory_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Memory percent",
		// 							"displayName": "Memory percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "iops",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "IO Operations per second",
		// 							"displayName": "IOPS",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage percent",
		// 							"displayName": "Storage percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_used",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage used",
		// 							"displayName": "Storage used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "active_connections",
		// 							"aggregationType": "Average",
		// 							"category": "Traffic",
		// 							"displayDescription": "Active Connections",
		// 							"displayName": "Active Connections",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "network_bytes_egress",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Network Out across active connections",
		// 							"displayName": "Network Out",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "network_bytes_ingress",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Network In across active connections",
		// 							"displayName": "Network In",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for PostgreSQL servers"),
		// 				Operation: to.Ptr("Read server log definitions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("The log definition of servers"),
		// 			},
		// 			Properties: map[string]any{
		// 				"serviceSpecification": map[string]any{
		// 					"logSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "PostgreSQLLogs",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "PostgreSQL Server Logs",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the disagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of servers or gets the properties for the specified server."),
		// 				Operation: to.Ptr("List/Get PostgreSQL Servers"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a server with the specified parameters or update the properties or tags for the specified server."),
		// 				Operation: to.Ptr("Create/Update PostgreSQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing server."),
		// 				Operation: to.Ptr("Delete PostgreSQL Server"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/firewallRules/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of firewall rules for a server or gets the properties for the specified firewall rule."),
		// 				Operation: to.Ptr("List/Get Firewall Rules"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/firewallRules/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates a firewall rule with the specified parameters or update an existing rule."),
		// 				Operation: to.Ptr("Create/Update Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/firewallRules/delete"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing firewall rule."),
		// 				Operation: to.Ptr("Delete Firewall Rule"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Firewall Rules"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/configurations/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return the list of configurations for a server or gets the properties for the specified configuration."),
		// 				Operation: to.Ptr("List/Get Configurations"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/configurations/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Update the value for the specified configuration"),
		// 				Operation: to.Ptr("Update Configuration"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Configurations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/serversv2/updateConfigurations/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Update configurations for the specified server"),
		// 				Operation: to.Ptr("Batch Update Server Configurations"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("PostgreSQL Server"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/providers/Microsoft.Insights/metricDefinitions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Return types of metrics that are available for databases"),
		// 				Operation: to.Ptr("Get database metric definitions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 			Properties: map[string]any{
		// 				"serviceSpecification": map[string]any{
		// 					"metricSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "backup_storage_used",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Backup Storage Used",
		// 							"displayName": "Backup Storage Used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "cpu_credits_consumed",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Total number of credits consumed by the database server",
		// 							"displayName": "CPU Credits Consumed",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "cpu_credits_remaining",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Total number of credits available to burst",
		// 							"displayName": "CPU Credits Remaining",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "cpu_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "CPU percent",
		// 							"displayName": "CPU percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "memory_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Memory percent",
		// 							"displayName": "Memory percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "iops",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "IO Operations per second",
		// 							"displayName": "IOPS",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "disk_queue_depth",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Number of outstanding I/O operations to the data disk",
		// 							"displayName": "Disk Queue Depth",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "read_throughput",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Bytes read per second from the data disk during monitoring period",
		// 							"displayName": "Read Throughput Bytes/Sec",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "write_throughput",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Bytes written per second to the data disk during monitoring period",
		// 							"displayName": "Write Throughput Bytes/Sec",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "read_iops",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Number of data disk I/O read operations per second",
		// 							"displayName": "Read IOPS",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "write_iops",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Number of data disk I/O write operations per second",
		// 							"displayName": "Write IOPS",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_percent",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage percent",
		// 							"displayName": "Storage percent",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Percent",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_used",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage used",
		// 							"displayName": "Storage used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "storage_free",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Storage Free",
		// 							"displayName": "Storage Free",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "txlogs_storage_used",
		// 							"aggregationType": "Average",
		// 							"category": "Saturation",
		// 							"displayDescription": "Transaction Log Storage Used",
		// 							"displayName": "Transaction Log Storage Used",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "active_connections",
		// 							"aggregationType": "Average",
		// 							"category": "Traffic",
		// 							"displayDescription": "Active Connections",
		// 							"displayName": "Active Connections",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "network_bytes_egress",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Network Out across active connections",
		// 							"displayName": "Network Out",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "network_bytes_ingress",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Network In across active connections",
		// 							"displayName": "Network In",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Bytes",
		// 						},
		// 						map[string]any{
		// 							"name": "connections_failed",
		// 							"aggregationType": "Total",
		// 							"category": "Errors",
		// 							"displayDescription": "Failed Connections",
		// 							"displayName": "Failed Connections",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "connections_succeeded",
		// 							"aggregationType": "Total",
		// 							"category": "Traffic",
		// 							"displayDescription": "Succeeded Connections",
		// 							"displayName": "Succeeded Connections",
		// 							"supportedAggregationTypes":[]any{
		// 								"Total",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 						map[string]any{
		// 							"name": "maximum_used_transactionIDs",
		// 							"aggregationType": "Average",
		// 							"category": "Traffic",
		// 							"displayDescription": "Maximum Used Transaction IDs",
		// 							"displayName": "Maximum Used Transaction IDs",
		// 							"supportedAggregationTypes":[]any{
		// 								"Average",
		// 								"Maximum",
		// 								"Minimum",
		// 							},
		// 							"unit": "Count",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/providers/Microsoft.Insights/logDefinitions/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the available logs for PostgreSQL servers"),
		// 				Operation: to.Ptr("Read server log definitions"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("The log definition of servers"),
		// 			},
		// 			Properties: map[string]any{
		// 				"serviceSpecification": map[string]any{
		// 					"logSpecifications":[]any{
		// 						map[string]any{
		// 							"name": "PostgreSQLLogs",
		// 							"blobDuration": "PT1H",
		// 							"displayName": "PostgreSQL Server Logs",
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/providers/Microsoft.Insights/diagnosticSettings/read"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Gets the disagnostic setting for the resource"),
		// 				Operation: to.Ptr("Read diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/providers/Microsoft.Insights/diagnosticSettings/write"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Creates or updates the diagnostic setting for the resource"),
		// 				Operation: to.Ptr("Write diagnostic setting"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Database Metric Definition"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateEndpointConnectionsApproval/action"),
		// 			Display: &armpostgresqlflexibleservers.OperationDisplay{
		// 				Description: to.Ptr("Determines if the user is allowed to approve a private endpoint connection"),
		// 				Operation: to.Ptr("Private Endpoint Connections Approval"),
		// 				Provider: to.Ptr("Microsoft DB for PostgreSQL"),
		// 				Resource: to.Ptr("Private Endpoint Connections Approval"),
		// 			},
		// 	}},
		// }
	}
}
