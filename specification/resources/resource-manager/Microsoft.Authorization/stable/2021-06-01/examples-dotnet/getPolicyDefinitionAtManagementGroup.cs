using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinitionAtManagementGroup.json
// this example is just showing the usage of "PolicyDefinitions_GetAtManagementGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupPolicyDefinitionResource created on azure
// for more information of creating ManagementGroupPolicyDefinitionResource, please refer to the document of ManagementGroupPolicyDefinitionResource
string managementGroupId = "MyManagementGroup";
string policyDefinitionName = "ResourceNaming";
ResourceIdentifier managementGroupPolicyDefinitionResourceId = ManagementGroupPolicyDefinitionResource.CreateResourceIdentifier(managementGroupId, policyDefinitionName);
ManagementGroupPolicyDefinitionResource managementGroupPolicyDefinition = client.GetManagementGroupPolicyDefinitionResource(managementGroupPolicyDefinitionResourceId);

// invoke the operation
ManagementGroupPolicyDefinitionResource result = await managementGroupPolicyDefinition.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
