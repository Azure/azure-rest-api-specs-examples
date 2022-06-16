import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/ListImageTemplates.json
     */
    /**
     * Sample code: List images by subscription.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void listImagesBySubscription(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().list(Context.NONE);
    }
}
