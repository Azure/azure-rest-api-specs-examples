using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/GetSecurityConnectorGovernanceRuleExecuteStatus_example.json
// this example is just showing the usage of "SecurityConnectorGovernanceRulesExecuteStatus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityConnectorGovernanceRuleResource created on azure
// for more information of creating SecurityConnectorGovernanceRuleResource, please refer to the document of SecurityConnectorGovernanceRuleResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "gcpResourceGroup";
string securityConnectorName = "gcpconnector";
string ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
ResourceIdentifier securityConnectorGovernanceRuleResourceId = SecurityConnectorGovernanceRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName, ruleId);
SecurityConnectorGovernanceRuleResource securityConnectorGovernanceRule = client.GetSecurityConnectorGovernanceRuleResource(securityConnectorGovernanceRuleResourceId);

// invoke the operation
string operationId = "58b33f4f-c8c7-4b01-99cc-d437db4d40dd";
ArmOperation<ExecuteRuleStatus> lro = await securityConnectorGovernanceRule.GetRuleExecutionStatusAsync(WaitUntil.Completed, operationId);
ExecuteRuleStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
