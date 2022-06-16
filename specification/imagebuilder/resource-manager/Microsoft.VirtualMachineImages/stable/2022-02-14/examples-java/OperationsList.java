import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/OperationsList.json
     */
    /**
     * Sample code: Retrieve operations list.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveOperationsList(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.operations().list(Context.NONE);
    }
}
