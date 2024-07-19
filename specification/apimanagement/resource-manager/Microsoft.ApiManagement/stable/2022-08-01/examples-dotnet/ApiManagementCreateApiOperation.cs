using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApiManagement.Models;
using Azure.ResourceManager.ApiManagement;

// Generated from example definition: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiOperation.json
// this example is just showing the usage of "ApiOperation_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApiResource created on azure
// for more information of creating ApiResource, please refer to the document of ApiResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string serviceName = "apimService1";
string apiId = "PetStoreTemplate2";
ResourceIdentifier apiResourceId = ApiResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, apiId);
ApiResource api = client.GetApiResource(apiResourceId);

// get the collection of this ApiOperationResource
ApiOperationCollection collection = api.GetApiOperations();

// invoke the operation
string operationId = "newoperations";
ApiOperationData data = new ApiOperationData()
{
    TemplateParameters =
    {
    },
    Description = "This can only be done by the logged in user.",
    Request = new RequestContract()
    {
        Description = "Created user object",
        QueryParameters =
        {
        },
        Headers =
        {
        },
        Representations =
        {
        new RepresentationContract("application/json")
        {
        SchemaId = "592f6c1d0af5840ca8897f0c",
        TypeName = "User",
        }
        },
    },
    Responses =
    {
    new ResponseContract(200)
    {
    Description = "successful operation",
    Representations =
    {
    new RepresentationContract("application/xml"),new RepresentationContract("application/json")
    },
    Headers =
    {
    },
    }
    },
    DisplayName = "createUser2",
    Method = "POST",
    UriTemplate = "/user1",
};
ArmOperation<ApiOperationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, operationId, data);
ApiOperationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApiOperationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
