using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Advisor;

// Generated from example definition: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateSuppression.json
// this example is just showing the usage of "Suppressions_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SuppressionContractResource created on azure
// for more information of creating SuppressionContractResource, please refer to the document of SuppressionContractResource
string resourceUri = "resourceUri";
string recommendationId = "recommendationId";
string name = "suppressionName1";
ResourceIdentifier suppressionContractResourceId = SuppressionContractResource.CreateResourceIdentifier(resourceUri, recommendationId, name);
SuppressionContractResource suppressionContract = client.GetSuppressionContractResource(suppressionContractResourceId);

// invoke the operation
SuppressionContractData data = new SuppressionContractData
{
    Ttl = "07:00:00:00",
};
ArmOperation<SuppressionContractResource> lro = await suppressionContract.UpdateAsync(WaitUntil.Completed, data);
SuppressionContractResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SuppressionContractData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
