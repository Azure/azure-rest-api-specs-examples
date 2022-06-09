```javascript
const { ServiceFabricManagementClient } = require("@azure/arm-servicefabric");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Service Fabric service resource with the specified name.
 *
 * @summary Create or update a Service Fabric service resource with the specified name.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ServicePutOperation_example_max.json
 */
async function putAServiceWithMaximumParameters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const serviceName = "myService";
  const parameters = {
    correlationScheme: [{ scheme: "Affinity", serviceName: "fabric:/app1/app1~svc1" }],
    defaultMoveCost: "Medium",
    partitionDescription: { partitionScheme: "Singleton" },
    placementConstraints: "NodeType==frontend",
    serviceDnsName: "my.service.dns",
    serviceKind: "Stateless",
    serviceLoadMetrics: [{ name: "metric1", weight: "Low" }],
    servicePackageActivationMode: "SharedProcess",
    servicePlacementPolicies: [],
    serviceTypeName: "myServiceType",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagementClient(credential, subscriptionId);
  const result = await client.services.beginCreateOrUpdateAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
    serviceName,
    parameters
  );
  console.log(result);
}

putAServiceWithMaximumParameters().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-servicefabric_2.0.1/sdk/servicefabric/arm-servicefabric/README.md) on how to add the SDK to your project and authenticate.
