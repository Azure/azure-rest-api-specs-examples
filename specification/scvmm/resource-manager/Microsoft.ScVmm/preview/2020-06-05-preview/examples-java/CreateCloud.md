```java
import com.azure.resourcemanager.scvmm.models.ExtendedLocation;

/** Samples for Clouds CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateCloud.json
     */
    /**
     * Sample code: CreateCloud.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void createCloud(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .clouds()
            .define("HRCloud")
            .withRegion("East US")
            .withExistingResourceGroup("testrg")
            .withExtendedLocation(
                new ExtendedLocation()
                    .withType("customLocation")
                    .withName(
                        "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"))
            .withUuid("aaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
            .withVmmServerId(
                "/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.SCVMM/VMMServers/ContosoVMMServer")
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-scvmm_1.0.0-beta.1/sdk/scvmm/azure-resourcemanager-scvmm/README.md) on how to add the SDK to your project and authenticate.
