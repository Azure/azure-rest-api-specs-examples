from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python replace_runbook_draft_content.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.runbook_draft.begin_replace_content(
        resource_group_name="rg",
        automation_account_name="ContoseAutomationAccount",
        runbook_name="Get-AzureVMTutorial",
        runbook_content="""<#
                .DESCRIPTION
                    An example runbook which prints out the first10 Azure VMs in your subscription (ordered alphabetically).
                    For more information about how this runbook authenticates to your Azure subscription, see our documentation here: http: //aka.ms/fxu3mn

                .NOTES
                    AUTHOR: Azure Automation Team
                    LASTEDIT: Mar27,
                    2015
            #>
            workflow Get-AzureVMTutorial{
                #The name of the Automation Credential Asset this runbook will use to authenticate to Azure.
                $CredentialAssetName = 'DefaultAzureCredential'

                #Get the credential with the above name from the Automation Asset store
                $Cred = Get-AutomationPSCredential -Name $CredentialAssetName
                if(!$Cred){
                    Throw"Could not find an Automation Credential Asset named '${CredentialAssetName}'. Make sure you have created one in this Automation Account."
                        }

                #Connect to your Azure Account
                $Account = Add-AzureAccount -Credential $Cred
                if(!$Account){
                    Throw"Could not authenticate to Azure using the credential asset '${CredentialAssetName}'. Make sure the user name and password are correct."
                        }

                #TODO (optional): pick the right subscription to use. Without this line, the default subscription for your Azure Account will be used.
                #Select-AzureSubscription -SubscriptionName"TODO: your Azure subscription name here"
                
                #Get all the VMs you have in your Azure subscription
                $VMs = Get-AzureVM

                #Print out up to10 of those VMs
                if(!$VMs){
                    Write-Output"No VMs were found in your subscription."
                        } else{
                    Write-Output $VMs[0..9
                            ]
                        }
                    }""",
    ).result()
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-08-08/examples/replaceRunbookDraftContent.json
if __name__ == "__main__":
    main()
