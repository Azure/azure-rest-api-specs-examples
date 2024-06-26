from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python create_api_polling.py

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
            "kind": "APIPolling",
            "properties": {
                "connectorUiConfig": {
                    "availability": {"isPreview": True, "status": 1},
                    "connectivityCriteria": [{"type": "SentinelKindsV2", "value": []}],
                    "dataTypes": [
                        {
                            "lastDataReceivedQuery": "{{graphQueriesTableName}}\n            | summarize Time = max(TimeGenerated)\n            | where isnotempty(Time)",
                            "name": "{{graphQueriesTableName}}",
                        }
                    ],
                    "descriptionMarkdown": "The GitHub audit log connector provides the capability to ingest GitHub logs into Azure Sentinel. By connecting GitHub audit logs into Azure Sentinel, you can view this data in workbooks, use it to create custom alerts, and improve your investigation process.",
                    "graphQueries": [
                        {
                            "baseQuery": "{{graphQueriesTableName}}",
                            "legend": "GitHub audit log events",
                            "metricName": "Total events received",
                        }
                    ],
                    "graphQueriesTableName": "GitHubAuditLogPolling_CL",
                    "instructionSteps": [
                        {
                            "description": "Enable GitHub audit Logs. \n Follow `this <https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/creating-a-personal-access-token>`_ to create or find your personal key",
                            "instructions": [
                                {
                                    "parameters": {
                                        "enable": "true",
                                        "userRequestPlaceHoldersInput": [
                                            {
                                                "displayText": "Organization Name",
                                                "placeHolderName": "{{placeHolder1}}",
                                                "placeHolderValue": "",
                                                "requestObjectKey": "apiEndpoint",
                                            }
                                        ],
                                    },
                                    "type": "APIKey",
                                }
                            ],
                            "title": "Connect GitHub Enterprise Audit Log to Azure Sentinel",
                        }
                    ],
                    "permissions": {
                        "customs": [
                            {
                                "description": "You need access to GitHub personal token, the key should have 'admin:org' scope",
                                "name": "GitHub API personal token Key",
                            }
                        ],
                        "resourceProvider": [
                            {
                                "permissionsDisplayText": "read and write permissions are required.",
                                "provider": "Microsoft.OperationalInsights/workspaces",
                                "providerDisplayName": "Workspace",
                                "requiredPermissions": {"delete": True, "read": True, "write": True},
                                "scope": "Workspace",
                            }
                        ],
                    },
                    "publisher": "GitHub",
                    "sampleQueries": [
                        {"description": "All logs", "query": "{{graphQueriesTableName}}\n | take 10 <change>"}
                    ],
                    "title": "GitHub Enterprise Audit Log",
                },
                "pollingConfig": {
                    "auth": {"apiKeyIdentifier": "token", "apiKeyName": "Authorization", "authType": "APIKey"},
                    "paging": {"pageSizeParaName": "per_page", "pagingType": "LinkHeader"},
                    "request": {
                        "apiEndpoint": "https://api.github.com/organizations/{{placeHolder1}}/audit-log",
                        "headers": {"Accept": "application/json", "User-Agent": "Scuba"},
                        "httpMethod": "Get",
                        "queryParameters": {"phrase": "created:{_QueryWindowStartTime}..{_QueryWindowEndTime}"},
                        "queryTimeFormat": "yyyy-MM-ddTHH:mm:ssZ",
                        "queryWindowInMin": 15,
                        "rateLimitQps": 50,
                        "retryCount": 2,
                        "timeoutInSeconds": 60,
                    },
                    "response": {"eventsJsonPaths": ["$"]},
                },
            },
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/dataConnectors/CreateAPIPolling.json
if __name__ == "__main__":
    main()
