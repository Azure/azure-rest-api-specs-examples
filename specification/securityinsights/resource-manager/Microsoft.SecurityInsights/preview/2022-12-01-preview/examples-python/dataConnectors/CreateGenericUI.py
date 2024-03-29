from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_generic_ui.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="d0cfe6b2-9ac0-4464-9919-dccaee2e48c0",
    )

    response = client.data_connectors.create_or_update(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        data_connector_id="316ec55e-7138-4d63-ab18-90c8a60fd1c8",
        data_connector={
            "kind": "GenericUI",
            "properties": {
                "connectorUiConfig": {
                    "availability": {"isPreview": True, "status": 1},
                    "connectivityCriteria": [
                        {
                            "type": "IsConnectedQuery",
                            "value": [
                                "{{graphQueriesTableName}}\n            | summarize LastLogReceived = max(TimeGenerated)\n            | project IsConnected = LastLogReceived > ago(30d)"
                            ],
                        }
                    ],
                    "dataTypes": [
                        {
                            "lastDataReceivedQuery": "{{graphQueriesTableName}}\n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)",
                            "name": "{{graphQueriesTableName}}",
                        }
                    ],
                    "descriptionMarkdown": "The [Qualys Vulnerability Management (VM)](https://www.qualys.com/apps/vulnerability-management/) data connector provides the capability to ingest vulnerability host detection data into Azure Sentinel through the Qualys API. The connector provides visibility into host detection data from vulerability scans. This connector provides Azure Sentinel the capability to view dashboards, create custom alerts, and improve investigation ",
                    "graphQueries": [
                        {
                            "baseQuery": "{{graphQueriesTableName}}",
                            "legend": "{{graphQueriesTableName}}",
                            "metricName": "Total data received",
                        }
                    ],
                    "graphQueriesTableName": "QualysHostDetection_CL",
                    "instructionSteps": [
                        {
                            "description": "..\n\n   **NOTE:** This connector uses Azure Functions to connect to Qualys VM to pull its logs into Azure Sentinel. This might result in additional data ingestion costs. Check the `Azure Functions pricing page <https://azure.microsoft.com/pricing/details/functions/>`_ for details.",
                            "title": "",
                        },
                        {
                            "description": "..\n\n   **(Optional Step)** Securely store workspace and API authorization key(s) or token(s) in Azure Key Vault. Azure Key Vault provides a secure mechanism to store and retrieve key values. `Follow these instructions <https://docs.microsoft.com/azure/app-service/app-service-key-vault-references>`_ to use Azure Key Vault with an Azure Function App.",
                            "title": "",
                        },
                        {
                            "description": "**STEP 1 - Configuration steps for the Qualys VM API**\n\n\n#. Log into the Qualys Vulnerability Management console with an administrator account, select the **Users** tab and the **Users** subtab. \n#. Click on the **New** drop-down menu and select **Users..**\n#. Create a username and password for the API account. \n#. In the **User Roles** tab, ensure the account role is set to **Manager** and access is allowed to **GUI** and **API**\n#. Log out of the administrator account and log into the console with the new API credentials for validation, then log out of the API account. \n#. Log back into the console using an administrator account and modify the API accounts User Roles, removing access to **GUI**. \n#. Save all changes.",
                            "title": "",
                        },
                        {
                            "description": "**STEP 2 - Choose ONE from the following two deployment options to deploy the connector and the associated Azure Function**\n\n..\n\n   **IMPORTANT:** Before deploying the Qualys VM connector, have the Workspace ID and Workspace Primary Key (can be copied from the following), as well as the Qualys VM API Authorization Key(s), readily available.",
                            "instructions": [
                                {
                                    "parameters": {"fillWith": ["WorkspaceId"], "label": "Workspace ID"},
                                    "type": "CopyableLabel",
                                },
                                {
                                    "parameters": {"fillWith": ["PrimaryKey"], "label": "Primary Key"},
                                    "type": "CopyableLabel",
                                },
                            ],
                            "title": "",
                        },
                        {
                            "description": 'Use this method for automated deployment of the Qualys VM connector using an ARM Tempate.\n\n\n#. \n   Click the **Deploy to Azure** button below. \n\n    \n   .. image:: https://aka.ms/deploytoazurebutton\n      :target: https://aka.ms/sentinelqualysvmazuredeploy\n      :alt: Deploy To Azure\n\n\n#. Select the preferred **Subscription**\\ , **Resource Group** and **Location**. \n#. Enter the **Workspace ID**\\ , **Workspace Key**\\ , **API Username**\\ , **API Password** , update the **URI**\\ , and any additional URI **Filter Parameters** (each filter should be separated by an "&" symbol, no spaces.) \n   ..\n\n      * Enter the URI that corresponds to your region. The complete list of API Server URLs can be `found here <https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf#G4.735348>`_ -- There is no need to add a time suffix to the URI, the Function App will dynamically append the Time Value to the URI in the proper format. \n      * The default **Time Interval** is set to pull the last five (5) minutes of data. If the time interval needs to be modified, it is recommended to change the Function App Timer Trigger accordingly (in the function.json file, post deployment) to prevent overlapping data ingestion. \n      * Note: If using Azure Key Vault secrets for any of the values above, use the\\ ``@Microsoft.KeyVault(SecretUri={Security Identifier})``\\ schema in place of the string values. Refer to `Key Vault references documentation <https://docs.microsoft.com/azure/app-service/app-service-key-vault-references>`_ for further details. \n\n\n#. Mark the checkbox labeled **I agree to the terms and conditions stated above**. \n#. Click **Purchase** to deploy.',
                            "title": "Option 1 - Azure Resource Manager (ARM) Template",
                        },
                        {
                            "description": "Use the following step-by-step instructions to deploy the Quayls VM connector manually with Azure Functions.",
                            "title": "Option 2 - Manual Deployment of Azure Functions",
                        },
                        {
                            "description": "**1. Create a Function App**\n\n\n#. From the Azure Portal, navigate to `Function App <https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.Web%2Fsites/kind/functionapp>`_\\ , and select **+ Add**.\n#. In the **Basics** tab, ensure Runtime stack is set to **Powershell Core**. \n#. In the **Hosting** tab, ensure the **Consumption (Serverless)** plan type is selected.\n#. Make other preferrable configuration changes, if needed, then click **Create**.",
                            "title": "",
                        },
                        {
                            "description": "**2. Import Function App Code**\n\n\n#. In the newly created Function App, select **Functions** on the left pane and click **+ New Function**.\n#. Select **Timer Trigger**.\n#. Enter a unique Function **Name** and leave the default cron schedule of every 5 minutes, then click **Create**.\n#. Click on **Code + Test** on the left pane. \n#. Copy the `Function App Code <https://aka.ms/sentinelqualysvmazurefunctioncode>`_ and paste into the Function App ``run.ps1`` editor.\n#. Click **Save**.",
                            "title": "",
                        },
                        {
                            "description": '**3. Configure the Function App**\n\n\n#. In the Function App, select the Function App Name and select **Configuration**.\n#. In the **Application settings** tab, select **+ New application setting**.\n#. Add each of the following seven (7) application settings individually, with their respective string values (case-sensitive): \n   .. code-block::\n\n       apiUsername\n       apiPassword\n       workspaceID\n       workspaceKey\n       uri\n       filterParameters\n       timeInterval\n\n   ..\n\n      * Enter the URI that corresponds to your region. The complete list of API Server URLs can be `found here <https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf#G4.735348>`_. The ``uri`` value must follow the following schema: ``https://<API Server>/api/2.0/fo/asset/host/vm/detection/?action=list&vm_processed_after=`` -- There is no need to add a time suffix to the URI, the Function App will dynamically append the Time Value to the URI in the proper format.\n      * Add any additional filter parameters, for the ``filterParameters`` variable, that need to be appended to the URI. Each parameter should be seperated by an "&" symbol and should not include any spaces.\n      * Set the ``timeInterval`` (in minutes) to the value of ``5`` to correspond to the Timer Trigger of every ``5`` minutes. If the time interval needs to be modified, it is recommended to change the Function App Timer Trigger accordingly to prevent overlapping data ingestion.\n      * Note: If using Azure Key Vault, use the\\ ``@Microsoft.KeyVault(SecretUri={Security Identifier})``\\ schema in place of the string values. Refer to `Key Vault references documentation <https://docs.microsoft.com/azure/app-service/app-service-key-vault-references>`_ for further details.\n\n\n#. Once all application settings have been entered, click **Save**.',
                            "title": "",
                        },
                        {
                            "description": '**4. Configure the host.json**.\n\nDue to the potentially large amount of Qualys host detection data being ingested, it can cause the execution time to surpass the default Function App timeout of five (5) minutes. Increase the default timeout duration to the maximum of ten (10) minutes, under the Consumption Plan, to allow more time for the Function App to execute.\n\n\n#. In the Function App, select the Function App Name and select the **App Service Editor** blade.\n#. Click **Go** to open the editor, then select the **host.json** file under the **wwwroot** directory.\n#. Add the line ``"functionTimeout": "00:10:00",`` above the ``managedDependancy`` line \n#. Ensure **SAVED** appears on the top right corner of the editor, then exit the editor.\n\n..\n\n   NOTE: If a longer timeout duration is required, consider upgrading to an `App Service Plan <https://docs.microsoft.com/azure/azure-functions/functions-scale#timeout>`_',
                            "title": "",
                        },
                    ],
                    "permissions": {
                        "customs": [
                            {
                                "description": "Read and write permissions to Azure Functions to create a Function App is required. `See the documentation to learn more about Azure Functions <https://docs.microsoft.com/azure/azure-functions/>`_.",
                                "name": "Microsoft.Web/sites permissions",
                            },
                            {
                                "description": "A Qualys VM API username and password is required. `See the documentation to learn more about Qualys VM API <https://www.qualys.com/docs/qualys-api-vmpc-user-guide.pdf>`_.",
                                "name": "Qualys API Key",
                            },
                        ],
                        "resourceProvider": [
                            {
                                "permissionsDisplayText": "read and write permissions on the workspace are required.",
                                "provider": "Microsoft.OperationalInsights/workspaces",
                                "providerDisplayName": "Workspace",
                                "requiredPermissions": {"delete": True, "read": True, "write": True},
                                "scope": "Workspace",
                            },
                            {
                                "permissionsDisplayText": "read permissions to shared keys for the workspace are required. [See the documentation to learn more about workspace keys](https://docs.microsoft.com/azure/azure-monitor/platform/agent-windows#obtain-workspace-id-and-key).",
                                "provider": "Microsoft.OperationalInsights/workspaces/sharedKeys",
                                "providerDisplayName": "Keys",
                                "requiredPermissions": {"action": True},
                                "scope": "Workspace",
                            },
                        ],
                    },
                    "publisher": "Qualys",
                    "sampleQueries": [
                        {
                            "description": "Top 10 Vulerabilities detected",
                            "query": "{{graphQueriesTableName}}\n | mv-expand todynamic(Detections_s)\n | extend Vulnerability = tostring(Detections_s.Results)\n | summarize count() by Vulnerability\n | top 10 by count_",
                        }
                    ],
                    "title": "Qualys Vulnerability Management (CCP DEMO)",
                }
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/dataConnectors/CreateGenericUI.json
if __name__ == "__main__":
    main()
