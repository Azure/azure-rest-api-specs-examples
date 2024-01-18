using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridContainerService;
using Azure.ResourceManager.HybridContainerService.Models;

// Generated from example definition: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/DeleteKubernetesVersions.json
// this example is just showing the usage of "DeleteKubernetesVersions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this KubernetesVersionProfileResource created on azure
// for more information of creating KubernetesVersionProfileResource, please refer to the document of KubernetesVersionProfileResource
string customLocationResourceUri = "subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation";
ResourceIdentifier kubernetesVersionProfileResourceId = KubernetesVersionProfileResource.CreateResourceIdentifier(customLocationResourceUri);
KubernetesVersionProfileResource kubernetesVersionProfile = client.GetKubernetesVersionProfileResource(kubernetesVersionProfileResourceId);

// invoke the operation
await kubernetesVersionProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
