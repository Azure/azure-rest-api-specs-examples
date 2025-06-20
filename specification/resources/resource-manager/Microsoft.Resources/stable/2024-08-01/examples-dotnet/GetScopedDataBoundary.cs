using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2024-08-01/examples/GetScopedDataBoundary.json
// this example is just showing the usage of "DataBoundaries_GetScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoundaryResource created on azure
// for more information of creating DataBoundaryResource, please refer to the document of DataBoundaryResource
string scope = "subscriptions/11111111-1111-1111-1111-111111111111/resourcegroups/my-resource-group";
DataBoundaryName name = DataBoundaryName.Default;
ResourceIdentifier dataBoundaryResourceId = DataBoundaryResource.CreateResourceIdentifier(scope, name);
DataBoundaryResource dataBoundary = client.GetDataBoundaryResource(dataBoundaryResourceId);

// invoke the operation
DataBoundaryName name0 = DataBoundaryName.Default;
DataBoundaryResource result = await dataBoundary.GetAsync(name0);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoundaryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
