using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Credentials_Get.json
// this example is just showing the usage of "CredentialOperations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryManagedIdentityCredentialResource created on azure
// for more information of creating DataFactoryManagedIdentityCredentialResource, please refer to the document of DataFactoryManagedIdentityCredentialResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string credentialName = "exampleCredential";
ResourceIdentifier dataFactoryManagedIdentityCredentialResourceId = DataFactoryManagedIdentityCredentialResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, credentialName);
DataFactoryManagedIdentityCredentialResource dataFactoryManagedIdentityCredential = client.GetDataFactoryManagedIdentityCredentialResource(dataFactoryManagedIdentityCredentialResourceId);

// invoke the operation
DataFactoryManagedIdentityCredentialResource result = await dataFactoryManagedIdentityCredential.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryManagedIdentityCredentialData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
