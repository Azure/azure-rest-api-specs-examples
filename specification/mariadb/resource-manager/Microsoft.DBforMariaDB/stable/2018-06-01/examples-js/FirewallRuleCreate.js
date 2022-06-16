const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function firewallRuleCreate() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "TestGroup";
  const serverName = "testserver";
  const firewallRuleName = "rule1";
  const parameters = {
    endIpAddress: "255.255.255.255",
    startIpAddress: "0.0.0.0",
  };
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
  const result = await client.firewallRules.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serverName,
    firewallRuleName,
    parameters
  );
  console.log(result);
}

firewallRuleCreate().catch(console.error);
