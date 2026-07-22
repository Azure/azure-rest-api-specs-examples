
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestApprovalStatus;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestResourceRequestAction;

/**
 * Samples for CommunityEndpoints HandleApprovalCreation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CommunityEndpoints_HandleApprovalCreation.json
     */
    /**
     * Sample code: CommunityEndpoints_HandleApprovalCreation.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        communityEndpointsHandleApprovalCreation(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communityEndpoints().handleApprovalCreation("rgopenapi", "TestMyCommunity", "TestMyCommunityEndpoint",
            new ApprovalCallbackRequest().withResourceRequestAction(ApprovalCallbackRequestResourceRequestAction.CREATE)
                .withApprovalStatus(ApprovalCallbackRequestApprovalStatus.APPROVED),
            com.azure.core.util.Context.NONE);
    }
}
