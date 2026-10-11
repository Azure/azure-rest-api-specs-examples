const { DataProtectionClient } = require("@azure/arm-dataprotection");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a backup instance in a backup vault
 *
 * @summary create or update a backup instance in a backup vault
 * x-ms-original-file: 2026-07-01/BackupInstanceOperations/PutBackupInstance_PostgreSqlFlexibleServerBackupDatasourceParameters.json
 */
async function createBackupInstanceWithPostgreSqlFlexibleServerBackupDatasourceParameters() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "62b829ee-7936-40c9-a1c9-47a93f9f3965";
  const client = new DataProtectionClient(credential, subscriptionId);
  const result = await client.backupInstances.createOrUpdate(
    "pgflexrg",
    "pgflexvault",
    "pgflexbi",
    {
      properties: {
        dataSourceInfo: {
          datasourceType: "Microsoft.DBforPostgreSQL/flexibleServers",
          objectType: "Datasource",
          resourceID:
            "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver",
          resourceLocation: "eastus2euap",
          resourceName: "pgflexserver",
          resourceType: "Microsoft.DBforPostgreSQL/flexibleServers",
          resourceUri:
            "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver",
        },
        dataSourceSetInfo: {
          datasourceType: "Microsoft.DBforPostgreSQL/flexibleServers",
          objectType: "DatasourceSet",
          resourceID:
            "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver",
          resourceLocation: "eastus2euap",
          resourceName: "pgflexserver",
          resourceType: "Microsoft.DBforPostgreSQL/flexibleServers",
          resourceUri:
            "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgflexserver",
        },
        friendlyName: "pgflexbi",
        objectType: "BackupInstance",
        policyInfo: {
          policyId:
            "/subscriptions/62b829ee-7936-40c9-a1c9-47a93f9f3965/resourceGroups/pgflexrg/providers/Microsoft.DataProtection/BackupVaults/pgflexvault/backupPolicies/pgflexpolicy",
          policyParameters: {
            backupDatasourceParametersList: [
              {
                objectType: "PostgreSqlFlexibleServerBackupDatasourceParameters",
                backupSolutionType: "PhysicalBackup",
              },
            ],
          },
        },
      },
    },
  );
  console.log(result);
}
