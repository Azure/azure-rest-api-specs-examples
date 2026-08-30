
import com.azure.resourcemanager.enclave.models.ApprovalDeletionCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalDeletionCallbackRequestResourceRequestAction;

/**
 * Samples for EnclaveEndpoints HandleApprovalDeletion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/EnclaveEndpoints_HandleApprovalDeletion.json
     */
    /**
     * Sample code: EnclaveEndpoints_HandleApprovalDeletion.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveEndpointsHandleApprovalDeletion(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveEndpoints().handleApprovalDeletion(
            "rgopenapi", "TestMyEnclave", "TestMyEnclaveEndpoint", new ApprovalDeletionCallbackRequest()
                .withResourceRequestAction(ApprovalDeletionCallbackRequestResourceRequestAction.CREATE),
            com.azure.core.util.Context.NONE);
    }
}
