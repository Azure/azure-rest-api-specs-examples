const { HDInsightManagementClient } = require("@azure/arm-hdinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new HDInsight cluster with the specified parameters.
 *
 * @summary Creates a new HDInsight cluster with the specified parameters.
 * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/CreateHDInsightClusterWithAvailabilityZones.json
 */
async function createClusterWithAvailabilityZones() {
  const subscriptionId = process.env["HDINSIGHT_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["HDINSIGHT_RESOURCE_GROUP"] || "rg1";
  const clusterName = "cluster1";
  const parameters = {
    properties: {
      clusterDefinition: {
        configurations: {
          "ambari-conf": {
            "database-name": "{ambari database name}",
            "database-server": "{sql server name}.database.windows.net",
            "database-user-name": "**********",
            "database-user-password": "**********",
          },
          gateway: {
            "restAuthCredential.isEnabled": true,
            "restAuthCredential.password": "**********",
            "restAuthCredential.username": "admin",
          },
          "hive-env": {
            hive_database: "Existing MSSQL Server database with SQL authentication",
            hive_database_name: "{hive metastore name}",
            hive_database_type: "mssql",
            hive_existing_mssql_server_database: "{hive metastore name}",
            hive_existing_mssql_server_host: "{sql server name}.database.windows.net",
            hive_hostname: "{sql server name}.database.windows.net",
          },
          "hive-site": {
            "javax.jdo.option.ConnectionDriverName": "com.microsoft.sqlserver.jdbc.SQLServerDriver",
            "javax.jdo.option.ConnectionPassword": "**********!",
            "javax.jdo.option.ConnectionURL":
              "jdbc:sqlserver://{sql server name}.database.windows.net;database={hive metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
            "javax.jdo.option.ConnectionUserName": "**********",
          },
          "oozie-env": {
            oozie_database: "Existing MSSQL Server database with SQL authentication",
            oozie_database_name: "{oozie metastore name}",
            oozie_database_type: "mssql",
            oozie_existing_mssql_server_database: "{oozie metastore name}",
            oozie_existing_mssql_server_host: "{sql server name}.database.windows.net",
            oozie_hostname: "{sql server name}.database.windows.net",
          },
          "oozie-site": {
            "oozie.db.schema.name": "oozie",
            "oozie.service.JPAService.jdbc.driver": "com.microsoft.sqlserver.jdbc.SQLServerDriver",
            "oozie.service.JPAService.jdbc.password": "**********",
            "oozie.service.JPAService.jdbc.url":
              "jdbc:sqlserver://{sql server name}.database.windows.net;database={oozie metastore name};encrypt=true;trustServerCertificate=true;create=false;loginTimeout=300;sendStringParametersAsUnicode=true;prepareSQL=0",
            "oozie.service.JPAService.jdbc.username": "**********",
          },
        },
        kind: "hadoop",
      },
      clusterVersion: "3.6",
      computeProfile: {
        roles: [
          {
            name: "headnode",
            hardwareProfile: { vmSize: "standard_d3" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
          {
            name: "workernode",
            hardwareProfile: { vmSize: "standard_d3" },
            osProfile: {
              linuxOperatingSystemProfile: {
                password: "**********",
                sshProfile: { publicKeys: [{ certificateData: "**********" }] },
                username: "sshuser",
              },
            },
            targetInstanceCount: 2,
            virtualNetworkProfile: {
              id: "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname",
              subnet:
                "/subscriptions/subId/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/vnetsubnet",
            },
          },
        ],
      },
      osType: "Linux",
      storageProfile: {
        storageaccounts: [
          {
            name: "mystorage",
            container: "containername",
            enableSecureChannel: true,
            isDefault: true,
            key: "storage account key",
          },
        ],
      },
    },
    zones: ["1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new HDInsightManagementClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateAndWait(
    resourceGroupName,
    clusterName,
    parameters
  );
  console.log(result);
}
