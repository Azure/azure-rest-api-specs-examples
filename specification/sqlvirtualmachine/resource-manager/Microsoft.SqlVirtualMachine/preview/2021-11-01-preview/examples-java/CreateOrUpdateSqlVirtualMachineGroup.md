Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-sqlvirtualmachine_1.0.0-beta.2/sdk/sqlvirtualmachine/azure-resourcemanager-sqlvirtualmachine/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.sqlvirtualmachine.models.SqlVmGroupImageSku;
import com.azure.resourcemanager.sqlvirtualmachine.models.WsfcDomainProfile;
import java.util.HashMap;
import java.util.Map;

/** Samples for SqlVirtualMachineGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/CreateOrUpdateSqlVirtualMachineGroup.json
     */
    /**
     * Sample code: Creates or updates a SQL virtual machine group.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void createsOrUpdatesASQLVirtualMachineGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager
            .sqlVirtualMachineGroups()
            .define("testvmgroup")
            .withRegion("northeurope")
            .withExistingResourceGroup("testrg")
            .withTags(mapOf("mytag", "myval"))
            .withSqlImageOffer("SQL2016-WS2016")
            .withSqlImageSku(SqlVmGroupImageSku.ENTERPRISE)
            .withWsfcDomainProfile(
                new WsfcDomainProfile()
                    .withDomainFqdn("testdomain.com")
                    .withOuPath("OU=WSCluster,DC=testdomain,DC=com")
                    .withClusterBootstrapAccount("testrpadmin")
                    .withClusterOperatorAccount("testrp@testdomain.com")
                    .withSqlServiceAccount("sqlservice@testdomain.com")
                    .withStorageAccountUrl("https://storgact.blob.core.windows.net/")
                    .withStorageAccountPrimaryKey("<primary storage access key>"))
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
```
