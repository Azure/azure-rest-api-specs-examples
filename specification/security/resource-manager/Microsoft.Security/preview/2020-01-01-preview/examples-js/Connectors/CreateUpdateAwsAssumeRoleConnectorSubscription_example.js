const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a cloud account connector or update an existing one. Connect to your cloud account. For AWS, use either account credentials or role-based authentication. For GCP, use account organization credentials.
 *
 * @summary Create a cloud account connector or update an existing one. Connect to your cloud account. For AWS, use either account credentials or role-based authentication. For GCP, use account organization credentials.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
 */
async function awsAssumeRoleCreateACloudAccountConnectorForASubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const connectorName = "aws_dev2";
  const connectorSetting = {
    authenticationDetails: {
      authenticationType: "awsAssumeRole",
      awsAssumeRoleArn: "arn:aws:iam::81231569658:role/AscConnector",
      awsExternalId: "20ff7fc3-e762-44dd-bd96-b71116dcdc23",
    },
    hybridComputeSettings: {
      autoProvision: "On",
      proxyServer: { ip: "167.220.197.140", port: "34" },
      region: "West US 2",
      resourceGroupName: "AwsConnectorRG",
      servicePrincipal: {
        applicationId: "ad9bcd79-be9c-45ab-abd8-80ca1654a7d1",
        secret: "<secret>",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.connectors.createOrUpdate(connectorName, connectorSetting);
  console.log(result);
}

awsAssumeRoleCreateACloudAccountConnectorForASubscription().catch(console.error);
