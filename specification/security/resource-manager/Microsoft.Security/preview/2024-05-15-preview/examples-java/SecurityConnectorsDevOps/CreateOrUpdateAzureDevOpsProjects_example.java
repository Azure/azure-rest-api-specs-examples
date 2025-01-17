
import com.azure.resourcemanager.security.models.ActionableRemediation;
import com.azure.resourcemanager.security.models.ActionableRemediationState;
import com.azure.resourcemanager.security.models.AzureDevOpsProjectProperties;
import com.azure.resourcemanager.security.models.OnboardingState;

/**
 * Samples for AzureDevOpsProjects CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/CreateOrUpdateAzureDevOpsProjects_example.json
     */
    /**
     * Sample code: CreateOrUpdate_AzureDevOpsProjects.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateAzureDevOpsProjects(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsProjects().define("myAzDevOpsProject")
            .withExistingAzureDevOpsOrg("myRg", "mySecurityConnectorName", "myAzDevOpsOrg")
            .withProperties(new AzureDevOpsProjectProperties().withOnboardingState(OnboardingState.NOT_APPLICABLE)
                .withActionableRemediation(new ActionableRemediation().withState(ActionableRemediationState.ENABLED)))
            .create();
    }
}
