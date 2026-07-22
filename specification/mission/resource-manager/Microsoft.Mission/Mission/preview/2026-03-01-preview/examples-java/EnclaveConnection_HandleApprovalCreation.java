
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestApprovalStatus;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestResourceRequestAction;

/**
 * Samples for EnclaveConnection HandleApprovalCreation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/EnclaveConnection_HandleApprovalCreation.json
     */
    /**
     * Sample code: EnclaveConnection_HandleApprovalCreation.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        enclaveConnectionHandleApprovalCreation(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.enclaveConnections().handleApprovalCreation("rgopenapi", "TestMyEnclaveConnection",
            new ApprovalCallbackRequest().withResourceRequestAction(ApprovalCallbackRequestResourceRequestAction.CREATE)
                .withApprovalStatus(ApprovalCallbackRequestApprovalStatus.APPROVED).withApprovalCallbackPayload(
                    "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}"),
            com.azure.core.util.Context.NONE);
    }
}
