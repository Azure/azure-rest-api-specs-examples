using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/contentPackages/GetPackageById.json
// this example is just showing the usage of "ContentPackages_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsPackageResource created on azure
// for more information of creating SecurityInsightsPackageResource, please refer to the document of SecurityInsightsPackageResource
string subscriptionId = "d0cfeab2-9ae0-4464-9919-dccaee2e48f0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string packageId = "str.azure-sentinel-solution-str";
ResourceIdentifier securityInsightsPackageResourceId = SecurityInsightsPackageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, packageId);
SecurityInsightsPackageResource securityInsightsPackage = client.GetSecurityInsightsPackageResource(securityInsightsPackageResourceId);

// invoke the operation
SecurityInsightsPackageResource result = await securityInsightsPackage.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityInsightsPackageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
