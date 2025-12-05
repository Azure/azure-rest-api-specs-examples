using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/IntegrationRuntimeObjectMetadata_Refresh.json
// this example is just showing the usage of "IntegrationRuntimeObjectMetadata_Refresh" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryIntegrationRuntimeResource created on azure
// for more information of creating DataFactoryIntegrationRuntimeResource, please refer to the document of DataFactoryIntegrationRuntimeResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string integrationRuntimeName = "testactivityv2";
ResourceIdentifier dataFactoryIntegrationRuntimeResourceId = DataFactoryIntegrationRuntimeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, integrationRuntimeName);
DataFactoryIntegrationRuntimeResource dataFactoryIntegrationRuntime = client.GetDataFactoryIntegrationRuntimeResource(dataFactoryIntegrationRuntimeResourceId);

// invoke the operation
ArmOperation<SsisObjectMetadataStatusResult> lro = await dataFactoryIntegrationRuntime.RefreshIntegrationRuntimeObjectMetadataAsync(WaitUntil.Completed);
SsisObjectMetadataStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
