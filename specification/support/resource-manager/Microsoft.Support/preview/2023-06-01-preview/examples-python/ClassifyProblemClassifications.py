from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.support import MicrosoftSupport

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-support
# USAGE
    python classify_problem_classifications.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MicrosoftSupport(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.problem_classifications_no_subscription.classify_problems(
        problem_service_name="serviceId1",
        problem_classifications_classification_input={"issueSummary": "Can not connect to Windows VM"},
    )
    print(response)


# x-ms-original-file: specification/support/resource-manager/Microsoft.Support/preview/2023-06-01-preview/examples/ClassifyProblemClassifications.json
if __name__ == "__main__":
    main()
