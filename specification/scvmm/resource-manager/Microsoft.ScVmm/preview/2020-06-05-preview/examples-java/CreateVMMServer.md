Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-scvmm_1.0.0-beta.1/sdk/scvmm/azure-resourcemanager-scvmm/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.scvmm.models.ExtendedLocation;
import com.azure.resourcemanager.scvmm.models.VmmServerPropertiesCredentials;

/** Samples for VmmServers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/scvmm/resource-manager/Microsoft.ScVmm/preview/2020-06-05-preview/examples/CreateVMMServer.json
     */
    /**
     * Sample code: CreateVMMServer.
     *
     * @param manager Entry point to ScvmmManager.
     */
    public static void createVMMServer(com.azure.resourcemanager.scvmm.ScvmmManager manager) {
        manager
            .vmmServers()
            .define("ContosoVMMServer")
            .withRegion("East US")
            .withExistingResourceGroup("testrg")
            .withExtendedLocation(
                new ExtendedLocation()
                    .withType("customLocation")
                    .withName(
                        "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso"))
            .withFqdn("VMM.contoso.com")
            .withCredentials(new VmmServerPropertiesCredentials().withUsername("testuser").withPassword("password"))
            .withPort(1234)
            .create();
    }
}
```
