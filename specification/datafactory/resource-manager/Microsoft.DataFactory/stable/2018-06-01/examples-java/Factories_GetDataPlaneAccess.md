Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.8/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.datafactory.models.UserAccessPolicy;

/** Samples for Factories GetDataPlaneAccess. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_GetDataPlaneAccess.json
     */
    /**
     * Sample code: Factories_GetDataPlaneAccess.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void factoriesGetDataPlaneAccess(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .factories()
            .getDataPlaneAccessWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                new UserAccessPolicy()
                    .withPermissions("r")
                    .withAccessResourcePath("")
                    .withProfileName("DefaultProfile")
                    .withStartTime("2018-11-10T02:46:20.2659347Z")
                    .withExpireTime("2018-11-10T09:46:20.2659347Z"),
                Context.NONE);
    }
}
```
