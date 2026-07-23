
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestApprovalStatus;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestResourceRequestAction;

/**
 * Samples for EnclaveEndpoints HandleApprovalCreation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EnclaveEndpoints_HandleApprovalCreation.json
     */
    /**
     * Sample code: EnclaveEndpoints_HandleApprovalCreation.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveEndpointsHandleApprovalCreation(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveEndpoints().handleApprovalCreation("rgopenapi", "TestMyEnclave", "TestMyEnclaveEndpoint",
            new ApprovalCallbackRequest().withResourceRequestAction(ApprovalCallbackRequestResourceRequestAction.CREATE)
                .withApprovalStatus(ApprovalCallbackRequestApprovalStatus.APPROVED),
            com.azure.core.util.Context.NONE);
    }
}
