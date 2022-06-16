import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.SaveImageBody;

/** Samples for LabPlans SaveImage. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/saveImageVirtualMachine.json
     */
    /**
     * Sample code: saveImageVirtualMachine.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void saveImageVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .labPlans()
            .saveImage(
                "testrg123",
                "testlabplan",
                new SaveImageBody()
                    .withName("Test Image")
                    .withLabVirtualMachineId(
                        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.LabServices/labs/testlab/virtualMachines/template"),
                Context.NONE);
    }
}
