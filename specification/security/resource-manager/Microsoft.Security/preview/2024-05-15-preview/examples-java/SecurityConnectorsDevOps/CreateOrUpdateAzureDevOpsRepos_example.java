
import com.azure.resourcemanager.security.models.ActionableRemediation;
import com.azure.resourcemanager.security.models.ActionableRemediationState;
import com.azure.resourcemanager.security.models.AzureDevOpsRepositoryProperties;
import com.azure.resourcemanager.security.models.OnboardingState;

/**
 * Samples for AzureDevOpsRepos CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-05-15-preview/examples/
     * SecurityConnectorsDevOps/CreateOrUpdateAzureDevOpsRepos_example.json
     */
    /**
     * Sample code: CreateOrUpdate_AzureDevOpsRepos.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateAzureDevOpsRepos(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.azureDevOpsRepos().define("myAzDevOpsRepo")
            .withExistingProject("myRg", "mySecurityConnectorName", "myAzDevOpsOrg", "myAzDevOpsProject")
            .withProperties(new AzureDevOpsRepositoryProperties().withOnboardingState(OnboardingState.NOT_APPLICABLE)
                .withActionableRemediation(new ActionableRemediation().withState(ActionableRemediationState.ENABLED)))
            .create();
    }
}
