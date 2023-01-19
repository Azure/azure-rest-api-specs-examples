using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDCustomDomains_RefreshValidationToken.json
// this example is just showing the usage of "FrontDoorCustomDomains_RefreshValidationToken" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorCustomDomainResource created on azure
// for more information of creating FrontDoorCustomDomainResource, please refer to the document of FrontDoorCustomDomainResource
string subscriptionId = "subid";
string resourceGroupName = "RG";
string profileName = "profile1";
string customDomainName = "domain1";
ResourceIdentifier frontDoorCustomDomainResourceId = FrontDoorCustomDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, profileName, customDomainName);
FrontDoorCustomDomainResource frontDoorCustomDomain = client.GetFrontDoorCustomDomainResource(frontDoorCustomDomainResourceId);

// invoke the operation
await frontDoorCustomDomain.RefreshValidationTokenAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
