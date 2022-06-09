```java
import com.azure.core.util.Context;

/** Samples for AzureBareMetalInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/stable/2021-08-09/examples/AzureBareMetalInstances_Get.json
     */
    /**
     * Sample code: Get an AzureBareMetal instance.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void getAnAzureBareMetalInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager
            .azureBareMetalInstances()
            .getByResourceGroupWithResponse("myResourceGroup", "myAzureBareMetalInstance", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-baremetalinfrastructure_1.0.0-beta.1/sdk/baremetalinfrastructure/azure-resourcemanager-baremetalinfrastructure/README.md) on how to add the SDK to your project and authenticate.
