import pytest

from tests.integrations import command as cmd


@pytest.mark.slow  # type: ignore
@pytest.mark.integ3  # type: ignore
def test_basic_notebook_start_and_kill() -> None:
    with cmd.interactive_command("notebook", "start") as notebook:
        for line in notebook.stdout:
            if "Jupyter Notebook is running" in line:
                return

    raise AssertionError()