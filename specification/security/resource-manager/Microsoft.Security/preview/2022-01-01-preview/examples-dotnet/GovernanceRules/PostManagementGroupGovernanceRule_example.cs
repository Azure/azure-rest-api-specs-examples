using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostManagementGroupGovernanceRule_example.json
// this example is just showing the usage of "GovernanceRules_Execute" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GovernanceRuleResource created on azure
// for more information of creating GovernanceRuleResource, please refer to the document of GovernanceRuleResource
string scope = "providers/Microsoft.Management/managementGroups/contoso";
string ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
ResourceIdentifier governanceRuleResourceId = GovernanceRuleResource.CreateResourceIdentifier(scope, ruleId);
GovernanceRuleResource governanceRule = client.GetGovernanceRuleResource(governanceRuleResourceId);

// invoke the operation
await governanceRule.ExecuteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
