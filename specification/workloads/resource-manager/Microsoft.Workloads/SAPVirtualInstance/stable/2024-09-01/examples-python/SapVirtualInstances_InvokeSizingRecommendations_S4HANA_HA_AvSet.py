from azure.identity import DefaultAzureCredential

from azure.mgmt.workloadssapvirtualinstance import WorkloadsSapVirtualInstanceMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-workloadssapvirtualinstance
# USAGE
    python sap_virtual_instances_invoke_sizing_recommendations_s4_hana_ha_av_set.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = WorkloadsSapVirtualInstanceMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.sap_virtual_instances.get_sizing_recommendations(
        location="centralus",
        body={
            "appLocation": "eastus",
            "databaseType": "HANA",
            "dbMemory": 1024,
            "dbScaleMethod": "ScaleUp",
            "deploymentType": "ThreeTier",
            "environment": "Prod",
            "highAvailabilityType": "AvailabilitySet",
            "sapProduct": "S4HANA",
            "saps": 75000,
        },
    )
    print(response)


# x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeSizingRecommendations_S4HANA_HA_AvSet.json
if __name__ == "__main__":
    main()
