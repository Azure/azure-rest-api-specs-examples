from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python put_metadata.py

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

    response = client.metadata.create(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        metadata_name="metadataName",
        metadata={
            "properties": {
                "author": {"email": "email@microsoft.com", "name": "User Name"},
                "categories": {"domains": ["Application", "Security – Insider Threat"], "verticals": ["Healthcare"]},
                "contentId": "c00ee137-7475-47c8-9cce-ec6f0f1bedd0",
                "contentSchemaVersion": "2.0",
                "customVersion": "1.0",
                "dependencies": {
                    "criteria": [
                        {
                            "criteria": [
                                {
                                    "contentId": "045d06d0-ee72-4794-aba4-cf5646e4c756",
                                    "kind": "DataConnector",
                                    "name": "Microsoft Defender for Endpoint",
                                },
                                {"contentId": "dbfcb2cc-d782-40ef-8d94-fe7af58a6f2d", "kind": "DataConnector"},
                                {
                                    "contentId": "de4dca9b-eb37-47d6-a56f-b8b06b261593",
                                    "kind": "DataConnector",
                                    "version": "2.0",
                                },
                            ],
                            "operator": "OR",
                        },
                        {"contentId": "31ee11cc-9989-4de8-b176-5e0ef5c4dbab", "kind": "Playbook", "version": "1.0"},
                        {"contentId": "21ba424a-9438-4444-953a-7059539a7a1b", "kind": "Parser"},
                    ],
                    "operator": "AND",
                },
                "firstPublishDate": "2021-05-18",
                "kind": "AnalyticsRule",
                "lastPublishDate": "2021-05-18",
                "parentId": "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName",
                "previewImages": ["firstImage.png", "secondImage.jpeg"],
                "previewImagesDark": ["firstImageDark.png", "secondImageDark.jpeg"],
                "providers": ["Amazon", "Microsoft"],
                "source": {
                    "kind": "Solution",
                    "name": "Contoso Solution 1.0",
                    "sourceId": "b688a130-76f4-4a07-bf57-762222a3cadf",
                },
                "support": {
                    "email": "support@microsoft.com",
                    "link": "https://support.microsoft.com/",
                    "name": "Microsoft",
                    "tier": "Partner",
                },
                "threatAnalysisTactics": ["reconnaissance", "commandandcontrol"],
                "threatAnalysisTechniques": ["T1548", "T1548.001"],
                "version": "1.0.0.0",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/metadata/PutMetadata.json
if __name__ == "__main__":
    main()
