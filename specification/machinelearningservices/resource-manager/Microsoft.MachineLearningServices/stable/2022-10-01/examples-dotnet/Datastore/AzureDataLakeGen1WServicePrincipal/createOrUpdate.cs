using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/Datastore/AzureDataLakeGen1WServicePrincipal/createOrUpdate.json
// this example is just showing the usage of "Datastores_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "my-aml-workspace";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningDatastoreResource
MachineLearningDatastoreCollection collection = machineLearningWorkspace.GetMachineLearningDatastores();

// invoke the operation
string name = "string";
MachineLearningDatastoreData data = new MachineLearningDatastoreData(new MachineLearningAzureDataLakeGen1Datastore(new MachineLearningServicePrincipalDatastoreCredentials(Guid.Parse("00000000-1111-2222-3333-444444444444"), new MachineLearningServicePrincipalDatastoreSecrets()
{
    ClientSecret = "string",
}, Guid.Parse("00000000-1111-2222-3333-444444444444"))
{
    AuthorityUri = new Uri("string"),
    ResourceUri = new Uri("string"),
}, "string")
{
    Description = "string",
    Properties =
    {
    },
    Tags =
    {
    ["string"] = "string",
    },
});
bool? skipValidation = false;
ArmOperation<MachineLearningDatastoreResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data, skipValidation: skipValidation);
MachineLearningDatastoreResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningDatastoreData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
