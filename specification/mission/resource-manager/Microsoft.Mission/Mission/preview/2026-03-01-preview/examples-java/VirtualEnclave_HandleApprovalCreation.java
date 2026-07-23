
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequest;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestApprovalStatus;
import com.azure.resourcemanager.enclave.models.ApprovalCallbackRequestResourceRequestAction;

/**
 * Samples for VirtualEnclave HandleApprovalCreation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/VirtualEnclave_HandleApprovalCreation.json
     */
    /**
     * Sample code: VirtualEnclave_HandleApprovalCreation.
     * 
     * @param manager Entry point to AzureEnclaveManager.
     */
    public static void
        virtualEnclaveHandleApprovalCreation(com.azure.resourcemanager.enclave.AzureEnclaveManager manager) {
        manager.virtualEnclaves().handleApprovalCreation("rgopenapi", "TestMyEnclave",
            new ApprovalCallbackRequest().withResourceRequestAction(ApprovalCallbackRequestResourceRequestAction.CREATE)
                .withApprovalStatus(ApprovalCallbackRequestApprovalStatus.APPROVED).withApprovalCallbackPayload(
                    "{\n  \"key1\": \"value1\",\n  \"key2\": \"value2\"\n}"),
            com.azure.core.util.Context.NONE);
    }
}
