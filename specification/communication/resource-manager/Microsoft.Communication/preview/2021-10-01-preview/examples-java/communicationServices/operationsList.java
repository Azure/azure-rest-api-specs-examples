import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/operationsList.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void operationsList(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.operations().list(Context.NONE);
    }
}
