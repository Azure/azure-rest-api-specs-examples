using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageMover.Models;
using Azure.ResourceManager.StorageMover;

// Generated from example definition: 2025-07-01/Endpoints_Delete.json
// this example is just showing the usage of "Endpoint_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageMoverEndpointResource created on azure
// for more information of creating StorageMoverEndpointResource, please refer to the document of StorageMoverEndpointResource
string subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
string resourceGroupName = "examples-rg";
string storageMoverName = "examples-storageMoverName";
string endpointName = "examples-endpointName";
ResourceIdentifier storageMoverEndpointResourceId = StorageMoverEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageMoverName, endpointName);
StorageMoverEndpointResource storageMoverEndpoint = client.GetStorageMoverEndpointResource(storageMoverEndpointResourceId);

// invoke the operation
await storageMoverEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
