
import com.azure.resourcemanager.security.models.ActionableRemediation;
import com.azure.resourcemanager.security.models.ActionableRemediationState;
import com.azure.resourcemanager.security.models.AzureDevOpsOrgProperties;
import com.azure.resourcemanager.security.models.OnboardingState;

/**
 * Samples for AzureDevOpsOrgs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * CreateOrUpdateAzureDevOpsOrgs_example.json
     */
    /**
     * Sample code: CreateOrUpdate_AzureDevOpsOrgs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateAzureDevOpsOrgs(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsOrgs().define("myAzDevOpsOrg")
            .withExistingSecurityConnector("myRg", "mySecurityConnectorName")
            .withProperties(new AzureDevOpsOrgProperties().withOnboardingState(OnboardingState.NOT_APPLICABLE)
                .withActionableRemediation(new ActionableRemediation().withState(ActionableRemediationState.ENABLED)))
            .create();
    }
}
