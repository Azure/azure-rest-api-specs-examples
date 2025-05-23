using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/PutSecurityConnectorApplication_example.json
// this example is just showing the usage of "SecurityConnectorApplications_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityConnectorApplicationResource created on azure
// for more information of creating SecurityConnectorApplicationResource, please refer to the document of SecurityConnectorApplicationResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "gcpResourceGroup";
string securityConnectorName = "gcpconnector";
string applicationId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
ResourceIdentifier securityConnectorApplicationResourceId = SecurityConnectorApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName, applicationId);
SecurityConnectorApplicationResource securityConnectorApplication = client.GetSecurityConnectorApplicationResource(securityConnectorApplicationResourceId);

// invoke the operation
SecurityApplicationData data = new SecurityApplicationData
{
    DisplayName = "GCP Admin's application",
    Description = "An application on critical GCP recommendations",
    SourceResourceType = ApplicationSourceResourceType.Assessments,
    ConditionSets = {BinaryData.FromObjectAsJson(new
    {
    conditions = new object[]
    {
    new Dictionary<string, object>
    {
    ["operator"] = "contains",
    ["property"] = "$.Id",
    ["value"] = "-prod-"
    }
    },
    })},
};
ArmOperation<SecurityConnectorApplicationResource> lro = await securityConnectorApplication.UpdateAsync(WaitUntil.Completed, data);
SecurityConnectorApplicationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
