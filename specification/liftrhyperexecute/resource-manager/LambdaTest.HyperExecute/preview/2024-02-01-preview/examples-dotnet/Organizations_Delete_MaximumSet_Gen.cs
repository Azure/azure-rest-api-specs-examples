using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.LambdaTestHyperExecute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.LambdaTestHyperExecute;

// Generated from example definition: 2024-02-01-preview/Organizations_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "OrganizationResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LambdaTestHyperExecuteOrganizationResource created on azure
// for more information of creating LambdaTestHyperExecuteOrganizationResource, please refer to the document of LambdaTestHyperExecuteOrganizationResource
string subscriptionId = "171E7A75-341B-4472-BC4C-7603C5AB9F32";
string resourceGroupName = "rgopenapi";
string organizationname = "testorg";
ResourceIdentifier lambdaTestHyperExecuteOrganizationResourceId = LambdaTestHyperExecuteOrganizationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, organizationname);
LambdaTestHyperExecuteOrganizationResource lambdaTestHyperExecuteOrganization = client.GetLambdaTestHyperExecuteOrganizationResource(lambdaTestHyperExecuteOrganizationResourceId);

// invoke the operation
await lambdaTestHyperExecuteOrganization.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
