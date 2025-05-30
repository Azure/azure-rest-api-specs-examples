using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HybridContainerService.Models;
using Azure.ResourceManager.HybridContainerService;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/GetVmSkus.json
// this example is just showing the usage of "GetVmSkus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridContainerServiceVmSkuResource created on azure
// for more information of creating HybridContainerServiceVmSkuResource, please refer to the document of HybridContainerServiceVmSkuResource
string customLocationResourceUri = "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation";
ResourceIdentifier hybridContainerServiceVmSkuResourceId = HybridContainerServiceVmSkuResource.CreateResourceIdentifier(customLocationResourceUri);
HybridContainerServiceVmSkuResource hybridContainerServiceVmSku = client.GetHybridContainerServiceVmSkuResource(hybridContainerServiceVmSkuResourceId);

// invoke the operation
HybridContainerServiceVmSkuResource result = await hybridContainerServiceVmSku.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HybridContainerServiceVmSkuData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
