using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HDInsight;
using Azure.ResourceManager.HDInsight.Models;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/DeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightPrivateEndpointConnectionResource created on azure
// for more information of creating HDInsightPrivateEndpointConnectionResource, please refer to the document of HDInsightPrivateEndpointConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string clusterName = "cluster1";
string privateEndpointConnectionName = "testprivateep.b3bf5fed-9b12-4560-b7d0-2abe1bba07e2";
ResourceIdentifier hdInsightPrivateEndpointConnectionResourceId = HDInsightPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, privateEndpointConnectionName);
HDInsightPrivateEndpointConnectionResource hdInsightPrivateEndpointConnection = client.GetHDInsightPrivateEndpointConnectionResource(hdInsightPrivateEndpointConnectionResourceId);

// invoke the operation
await hdInsightPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
