using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/GlobalParameters_Create.json
// this example is just showing the usage of "GlobalParameters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryResource created on azure
// for more information of creating DataFactoryResource, please refer to the document of DataFactoryResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
ResourceIdentifier dataFactoryResourceId = DataFactoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName);
DataFactoryResource dataFactory = client.GetDataFactoryResource(dataFactoryResourceId);

// get the collection of this DataFactoryGlobalParameterResource
DataFactoryGlobalParameterCollection collection = dataFactory.GetDataFactoryGlobalParameters();

// invoke the operation
string globalParameterName = "default";
DataFactoryGlobalParameterData data = new DataFactoryGlobalParameterData(new Dictionary<string, DataFactoryGlobalParameterProperties>
{
    ["waitTime"] = new DataFactoryGlobalParameterProperties(DataFactoryGlobalParameterType.Int, BinaryData.FromObjectAsJson("5"))
});
ArmOperation<DataFactoryGlobalParameterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, globalParameterName, data);
DataFactoryGlobalParameterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryGlobalParameterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
