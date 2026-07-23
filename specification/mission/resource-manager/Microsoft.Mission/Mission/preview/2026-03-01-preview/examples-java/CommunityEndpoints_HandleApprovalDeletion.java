
import com.azure.resourcemanager.enclave.models.ApprovalDeletionCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalDeletionCallbackRequestResourceRequestAction;

/**
 * Samples for CommunityEndpoints HandleApprovalDeletion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/CommunityEndpoints_HandleApprovalDeletion.json
     */
    /**
     * Sample code: CommunityEndpoints_HandleApprovalDeletion.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        communityEndpointsHandleApprovalDeletion(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.communityEndpoints().handleApprovalDeletion(
            "rgopenapi", "TestMyCommunity", "TestMyCommunityEndpoint", new ApprovalDeletionCallbackRequest()
                .withResourceRequestAction(ApprovalDeletionCallbackRequestResourceRequestAction.CREATE),
            com.azure.core.util.Context.NONE);
    }
}
