<a href="https://multiparty.ai"><img alt="Multiparty logo" style="height:48px" src="https://github.com/indentapis/multiparty/assets/1026125/c99e043a-df04-46d9-b269-6dd90e927aa0" /></a>

Multiparty provides fine-grained, programmatically defined permissions to control sensitive operations performed by AI Agents, without brittle if statements.

[Sign up for the waitlist](https://multiparty.ai/) to give feedback and get access.

This repository serves as a central hub for the GitHub issues, use cases, and bug reports related to Multiparty for easier tracking and communication.
For additional information, visit [Multiparty.ai](https://multiparty.ai/).

#### Learn more

- [What is it?](https://multiparty.ai/#about)
- [Join the party](https://multiparty.ai/)
- [Playground](https://multiparty.ai/playground)

#### Feedback?

We'd love to hear from you! Feel free to open an issue about a new use case or [contact our team over email](mailto:open@indent.com).

## Concepts

- **Rules** prescribe expectations for a system
  - Deterministic evaluation using Google's Common Expression Language (CEL)
- **Actions** run using WASM or Linux containers
- **Plans** specify ways to satisfy Rules using Actions

## Architecture

### 🐚 Client

- Lightweight language wrappers around Rust/Go core
  - Python
  - JavaScript/TypeScript
  - Swift
  - Ruby
- Implements `check` and `enforce` commands
  - `check` quickly verifies assertions, e.g. don't execute if over budget
  - `enforce` provides async verification over longer windows, e.g. require user approval to be collected
- Support for local/embedded or remote engine
  - Embedded engine allows `check` without network connection
  - Remote engine provides centralized manage and control of rules

### ⚙️ Engine

- Evalute **Rules** either locally or remote
  - Local for edge compute, testing, and single-tenant applications
  - Remote allows distributed and remote coordination
- Perform **Actions** in accordance to Plans
  - Dispatches WASM or containers for execution
  - Capture results and side effects to be used by Rules
- Durable storage for **Plan** and `enforce` calls
  - Execute Actions using either WASM or containers
- Signed decision chains skip the guesswork on why an Action was taken
  - Audit logs satisfy compliance requirements
- Budgets allow LLM-backed applications to deliver high quality UX with scalable costs

## Frequently Asked Questions

<details>
  <summary><b>Where does this fit into my existing Langchain, LlamaIndex, or Griptape app?</b></summary>

```diff
# langchain/libs/langchain/langchain/callbacks/human.py

from typing import Any, Callable, Dict, Optional
from uuid import UUID

from langchain.callbacks.base import BaseCallbackHandler
+ from multiparty.utils import async_input

def _default_approve(_input: str) -> bool:
    msg = (
        "Do you approve of the following input? "
-        "Anything except 'Y'/'Yes' (case-insensitive) will be treated as a no."
    )
    msg += "\n\n" + _input + "\n"
-   resp = input(msg)
-   return resp.lower() in ("yes", "y")
+   return async_input(msg)


def _default_true(_: Dict[str, Any]) -> bool:
    return True


class HumanRejectedException(Exception):
    """Exception to raise when a person manually review and rejects a value."""


class HumanApprovalCallbackHandler(BaseCallbackHandler):
    """Callback for manually validating values."""

    raise_error: bool = True

    def __init__(
        self,
        approve: Callable[[Any], bool] = _default_approve,
        should_check: Callable[[Dict[str, Any]], bool] = _default_true,
    ):
        self._approve = approve
        self._should_check = should_check

    def on_tool_start(
        self,
        serialized: Dict[str, Any],
        input_str: str,
        *,
        run_id: UUID,
        parent_run_id: Optional[UUID] = None,
        **kwargs: Any,
    ) -> Any:
        if self._should_check(serialized) and not self._approve(input_str):
            raise HumanRejectedException(
                f"Inputs {input_str} to tool {serialized} were rejected."
            )
```

</details>

<details>
  <summary><b>How does this compare to existing rule engines?</b></summary>

<br />

**Rules** 

The Common Expression Language (CEL) is a non-Turing complete language designed
for simplicity, speed, safety, and portability. CEL's C-like [syntax][1] looks
nearly identical to equivalent expressions in C++, Go, Java, and TypeScript.

```java
// Check whether a resource name starts with a group name.
resource.name.startsWith("/groups/" + auth.claims.group)
```

```go
// Determine whether the request is in the permitted time window.
request.time - resource.age < duration("24h")
```

```typescript
// Check whether all resource names in a list match a given filter.
auth.claims.email_verified && resources.all(r, r.startsWith(auth.claims.email))
```

A CEL "program" is a single expression. The examples have been tagged as
`java`, `go`, and `typescript` within the markdown to showcase the commonality
of the syntax.

**Actions**

Calling external functions works out-of-the-box with multiple languages (Python, JavaScript, etc) with built-in caching. Read or write from any HTTP source.

**Tests**

Deploy automated systems with confidence that common sense controls are correctly implemented and easily tested.

</details>
