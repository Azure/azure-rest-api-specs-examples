```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.redisenterprise.models.AccessKeyType;
import com.azure.resourcemanager.redisenterprise.models.RegenerateKeyParameters;

/** Samples for Databases RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2022-01-01/examples/RedisEnterpriseDatabasesRegenerateKey.json
     */
    /**
     * Sample code: RedisEnterpriseDatabasesRegenerateKey.
     *
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseDatabasesRegenerateKey(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager
            .databases()
            .regenerateKey(
                "rg1",
                "cache1",
                "default",
                new RegenerateKeyParameters().withKeyType(AccessKeyType.PRIMARY),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-redisenterprise_1.1.0-beta.1/sdk/redisenterprise/azure-resourcemanager-redisenterprise/README.md) on how to add the SDK to your project and authenticate.
