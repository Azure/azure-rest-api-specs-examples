
import com.azure.resourcemanager.security.models.ActionableRemediation;
import com.azure.resourcemanager.security.models.ActionableRemediationState;
import com.azure.resourcemanager.security.models.AzureDevOpsRepository;
import com.azure.resourcemanager.security.models.AzureDevOpsRepositoryProperties;
import com.azure.resourcemanager.security.models.OnboardingState;

/**
 * Samples for AzureDevOpsRepos Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-04-01/examples/SecurityConnectorsDevOps/
     * UpdateAzureDevOpsRepos_example.json
     */
    /**
     * Sample code: Update_AzureDevOpsRepos.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void updateAzureDevOpsRepos(com.azure.resourcemanager.security.SecurityManager manager) {
        AzureDevOpsRepository resource = manager.azureDevOpsRepos().getWithResponse("myRg", "mySecurityConnectorName",
            "myAzDevOpsOrg", "myAzDevOpsProject", "myAzDevOpsRepo", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new AzureDevOpsRepositoryProperties().withOnboardingState(OnboardingState.NOT_APPLICABLE)
                .withActionableRemediation(new ActionableRemediation().withState(ActionableRemediationState.ENABLED)))
            .apply();
    }
}
