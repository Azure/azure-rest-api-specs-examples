using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedServiceIdentities;

// Generated from example definition: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/FederatedIdentityCredentialDelete.json
// this example is just showing the usage of "FederatedIdentityCredentials_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FederatedIdentityCredentialResource created on azure
// for more information of creating FederatedIdentityCredentialResource, please refer to the document of FederatedIdentityCredentialResource
string subscriptionId = "c267c0e7-0a73-4789-9e17-d26aeb0904e5";
string resourceGroupName = "rgName";
string resourceName = "resourceName";
string federatedIdentityCredentialResourceName = "ficResourceName";
ResourceIdentifier federatedIdentityCredentialResourceId = FederatedIdentityCredentialResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, federatedIdentityCredentialResourceName);
FederatedIdentityCredentialResource federatedIdentityCredential = client.GetFederatedIdentityCredentialResource(federatedIdentityCredentialResourceId);

// invoke the operation
await federatedIdentityCredential.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
