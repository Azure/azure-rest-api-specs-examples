using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConnectedVMwarevSphere;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteDatastore.json
// this example is just showing the usage of "Datastores_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VMwareDatastoreResource created on azure
// for more information of creating VMwareDatastoreResource, please refer to the document of VMwareDatastoreResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string datastoreName = "HRDatastore";
ResourceIdentifier vMwareDatastoreResourceId = VMwareDatastoreResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, datastoreName);
VMwareDatastoreResource vMwareDatastore = client.GetVMwareDatastoreResource(vMwareDatastoreResourceId);

// invoke the operation
await vMwareDatastore.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
