using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/Credentials_Create.json
// this example is just showing the usage of "CredentialOperations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryServiceCredentialResource created on azure
// for more information of creating DataFactoryServiceCredentialResource, please refer to the document of DataFactoryServiceCredentialResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string credentialName = "exampleCredential";
ResourceIdentifier dataFactoryServiceCredentialResourceId = DataFactoryServiceCredentialResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, credentialName);
DataFactoryServiceCredentialResource dataFactoryServiceCredential = client.GetDataFactoryServiceCredentialResource(dataFactoryServiceCredentialResourceId);

// invoke the operation
DataFactoryServiceCredentialData data = new DataFactoryServiceCredentialData(new DataFactoryManagedIdentityCredentialProperties
{
    ResourceId = new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-12345678abc/resourcegroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleUami"),
});
ArmOperation<DataFactoryServiceCredentialResource> lro = await dataFactoryServiceCredential.UpdateAsync(WaitUntil.Completed, data);
DataFactoryServiceCredentialResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryServiceCredentialData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
