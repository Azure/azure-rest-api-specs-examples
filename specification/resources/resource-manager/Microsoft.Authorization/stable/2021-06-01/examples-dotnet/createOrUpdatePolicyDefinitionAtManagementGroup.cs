using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAtManagementGroup.json
// this example is just showing the usage of "PolicyDefinitions_CreateOrUpdateAtManagementGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "MyManagementGroup";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// get the collection of this ManagementGroupPolicyDefinitionResource
ManagementGroupPolicyDefinitionCollection collection = managementGroupResource.GetManagementGroupPolicyDefinitions();

// invoke the operation
string policyDefinitionName = "ResourceNaming";
PolicyDefinitionData data = new PolicyDefinitionData()
{
    Mode = "All",
    DisplayName = "Enforce resource naming convention",
    Description = "Force resource names to begin with given 'prefix' and/or end with given 'suffix'",
    PolicyRule = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["if"] = new Dictionary<string, object>()
        {
            ["not"] = new Dictionary<string, object>()
            {
                ["field"] = "name",
                ["like"] = "[concat(parameters('prefix'), '*', parameters('suffix'))]"
            }
        },
        ["then"] = new Dictionary<string, object>()
        {
            ["effect"] = "deny"
        }
    }),
    Metadata = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["category"] = "Naming"
    }),
    Parameters =
    {
    ["prefix"] = new ArmPolicyParameter()
    {
    ParameterType = ArmPolicyParameterType.String,
    Metadata = new ParameterDefinitionsValueMetadata()
    {
    DisplayName = "Prefix",
    Description = "Resource name prefix",
    },
    },
    ["suffix"] = new ArmPolicyParameter()
    {
    ParameterType = ArmPolicyParameterType.String,
    Metadata = new ParameterDefinitionsValueMetadata()
    {
    DisplayName = "Suffix",
    Description = "Resource name suffix",
    },
    },
    },
};
ArmOperation<ManagementGroupPolicyDefinitionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyDefinitionName, data);
ManagementGroupPolicyDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
