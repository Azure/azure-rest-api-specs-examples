const { DevSpacesManagementClient } = require("@azure/arm-devspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing Azure Dev Spaces Controller with the specified update parameters.
 *
 * @summary Updates the properties of an existing Azure Dev Spaces Controller with the specified update parameters.
 * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ControllersUpdate_example.json
 */
async function controllersUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const name = "myControllerResource";
  const controllerUpdateParameters = {
    tags: { key: "value" },
    targetContainerHostCredentialsBase64: "QmFzZTY0IEVuY29kZWQgVmFsdWUK",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevSpacesManagementClient(credential, subscriptionId);
  const result = await client.controllers.update(
    resourceGroupName,
    name,
    controllerUpdateParameters
  );
  console.log(result);
}

controllersUpdate().catch(console.error);
