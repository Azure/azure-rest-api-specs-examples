
import com.azure.resourcemanager.enclave.models.ApprovalDeletionCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalDeletionCallbackRequestResourceRequestAction;

/**
 * Samples for EnclaveConnection HandleApprovalDeletion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EnclaveConnection_HandleApprovalDeletion.json
     */
    /**
     * Sample code: EnclaveConnection_HandleApprovalDeletion.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveConnectionHandleApprovalDeletion(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveConnections()
            .handleApprovalDeletion("rgopenapi", "TestMyEnclaveConnection", new ApprovalDeletionCallbackRequest()
                .withResourceRequestAction(ApprovalDeletionCallbackRequestResourceRequestAction.CREATE),
                com.azure.core.util.Context.NONE);
    }
}
